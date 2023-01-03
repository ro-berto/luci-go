// Copyright 2022 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package analysis

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"

	"go.chromium.org/luci/analysis/internal/analysis/metrics"
	"go.chromium.org/luci/analysis/internal/bqutil"
	"go.chromium.org/luci/analysis/internal/clustering"
	"go.chromium.org/luci/analysis/internal/clustering/algorithms/rulesalgorithm"
	configpb "go.chromium.org/luci/analysis/proto/config"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/trace"
)

// Cluster contains detailed information about a cluster, including
// a statistical summary of a cluster's failures, and their impact.
type Cluster struct {
	ClusterID clustering.ClusterID

	// MetricValues the values of cluster metrics. Only metrics which
	// have been computed for the cluster are populated.
	MetricValues map[metrics.ID]metrics.TimewiseCounts

	// The number of distinct user (i.e not automation generated) CLs
	// which have failures that are part of this cluster, over the last
	// 7 days. If this is more than a couple, it is a good indicator the
	// problem is really in the tree and not only on a few unsubmitted CLs.
	DistinctUserCLsWithFailures7d metrics.Counts
	// The number of postsubmit builds which have failures that are
	// a part of this cluster. If this is non-zero, it is an indicator
	// the problem is in the tree and not in a few unsubmitted CLs.
	PostsubmitBuildsWithFailures7d metrics.Counts

	// The realm(s) examples of the cluster are present in.
	Realms               []string
	ExampleFailureReason bigquery.NullString
	// Top Test IDs included in the cluster, up to 5. Unless the cluster
	// is empty, will always include at least one Test ID.
	TopTestIDs []TopCount
	// Top Monorail Components indicates the top monorail components failures
	// in the cluster are associated with by number of failures, up to 5.
	TopMonorailComponents []TopCount
}

// ExampleTestID returns an example Test ID that is part of the cluster, or
// "" if the cluster is empty.
func (s *Cluster) ExampleTestID() string {
	if len(s.TopTestIDs) > 0 {
		return s.TopTestIDs[0].Value
	}
	return ""
}

// TopCount captures the result of the APPROX_TOP_COUNT operator. See:
// https://cloud.google.com/bigquery/docs/reference/standard-sql/approximate_aggregate_functions#approx_top_count
type TopCount struct {
	// Value is the value that was frequently occurring.
	Value string
	// Count is the frequency with which the value occurred.
	Count int64
}

// RebuildAnalysis re-builds the cluster summaries analysis from
// clustered test results.
func (c *Client) RebuildAnalysis(ctx context.Context, luciProject string) error {
	datasetID, err := bqutil.DatasetForProject(luciProject)
	if err != nil {
		return errors.Annotate(err, "getting dataset").Err()
	}
	dataset := c.client.Dataset(datasetID)

	dstTable := dataset.Table("cluster_summaries")

	var precomputeList []string
	selectLists := make([][]string, len(metrics.CalculationBases))
	for _, metric := range metrics.ComputedMetrics {
		metricFilterIdentifier := "TRUE"
		if metric.FilterSQL != "" {
			metricFilterIdentifier = metric.ColumnName("filter")
			precomputeList = append(precomputeList, fmt.Sprintf("%s AS %s,", metric.FilterSQL, metricFilterIdentifier))
		}
		var itemIdentifier string
		if metric.CountSQL != "" {
			itemIdentifier = metric.ColumnName("item")
			precomputeList = append(precomputeList, fmt.Sprintf("%s AS %s,", metric.CountSQL, itemIdentifier))
		}

		for i, calculationBasis := range metrics.CalculationBases {
			// Further restrict to last 1/3/7 days.
			furtherFilter := fmt.Sprintf("is_%vd", calculationBasis.IntervalDays)
			if calculationBasis.Residual {
				furtherFilter += " AND f.is_included_with_high_priority"
			}
			var metricExpr string
			if metric.CountSQL != "" {
				metricExpr = fmt.Sprintf("COUNT(DISTINCT IF(%s AND %s, %s, NULL))", furtherFilter, metricFilterIdentifier, itemIdentifier)
			} else {
				metricExpr = fmt.Sprintf("COUNTIF(%s AND %s)", furtherFilter, metricFilterIdentifier)
			}
			// Each permutation will produce one JSON column,
			// inside of which there will be one field per metric.
			selectLists[i] = append(selectLists[i], fmt.Sprintf("%s AS %s", metricExpr, metric.BaseColumnName))
		}
	}
	var selectList []string
	for i, calculationBasis := range metrics.CalculationBases {
		columnName := fmt.Sprintf("metrics_%s", calculationBasis.ColumnSuffix())
		expr := `TO_JSON(STRUCT(` + strings.Join(selectLists[i], ",\n") + `))`
		selectList = append(selectList, fmt.Sprintf("%s AS %s,", expr, columnName))
	}

	q := c.client.Query(`
		WITH clustered_failures_latest AS (
		  SELECT
			cluster_algorithm,
			cluster_id,
			test_result_system,
			test_result_id,
			partition_time,
			ARRAY_AGG(cf ORDER BY last_updated DESC LIMIT 1)[OFFSET(0)] as f
		  FROM clustered_failures cf
		  WHERE partition_time >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 7 DAY)
		  GROUP BY cluster_algorithm, cluster_id, test_result_system, test_result_id, partition_time
		  HAVING f.is_included
		),
		clustered_failures_precompute AS (
		  SELECT
			cluster_algorithm,
			cluster_id,
			f,
			f.partition_time >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 1 DAY) as is_1d,
			f.partition_time >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 3 DAY) as is_3d,
			f.partition_time >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 7 DAY) as is_7d,
			` + strings.Join(precomputeList, "\n") + `
			-- The identity of the first changelist that was tested, assuming the
			-- result was part of a presubmit run, and the owner of the presubmit
			-- run was a user and not automation. Note SAFE_OFFSET runs NULL
			-- if the item does not exist, and CONCAT returns NULL if any argument
			-- is null.
			IF(f.presubmit_run_owner='user',
				CONCAT(f.changelists[SAFE_OFFSET(0)].host, '/', f.changelists[SAFE_OFFSET(0)].change),
				NULL) as presubmit_run_user_cl_id,
			f.changelists IS NULL OR ARRAY_LENGTH(f.changelists) = 0 AS is_postsubmit,
		  FROM clustered_failures_latest
		)
		SELECT
			cluster_algorithm,
			cluster_id,
			` + strings.Join(selectList, "\n") + `

			-- Analysis of whether the cluster occurs within the tree or only in isolated CLs.
			-- TODO(b/260631527): Determine whether we still need these metrics.
			COUNT(DISTINCT IF(is_7d, presubmit_run_user_cl_id, NULL)) as distinct_user_cls_with_failures_7d,
			COUNT(DISTINCT IF(is_7d AND is_postsubmit, f.ingested_invocation_id, NULL)) as postsubmit_builds_with_failures_7d,
			COUNT(DISTINCT IF(is_7d AND f.is_included_with_high_priority, presubmit_run_user_cl_id, NULL)) as distinct_user_cls_with_failures_residual_7d,
			COUNT(DISTINCT IF(is_7d AND is_postsubmit AND f.is_included_with_high_priority, f.ingested_invocation_id, NULL)) as postsubmit_builds_with_failures_residual_7d,

			-- Other analysis.
			ANY_VALUE(f.failure_reason) as example_failure_reason,
			ARRAY_AGG(DISTINCT f.realm) as realms,
			APPROX_TOP_COUNT(f.test_id, 5) as top_test_ids,
			APPROX_TOP_COUNT(IF(f.bug_tracking_component.system = 'monorail', f.bug_tracking_component.component, NULL), 5) as top_monorail_components,
		FROM clustered_failures_precompute
		GROUP BY cluster_algorithm, cluster_id
	`)
	q.DefaultDatasetID = dataset.DatasetID
	q.Dst = dstTable
	q.CreateDisposition = bigquery.CreateIfNeeded
	q.WriteDisposition = bigquery.WriteTruncate
	job, err := q.Run(ctx)
	if err != nil {
		return errors.Annotate(err, "starting cluster summary analysis").Err()
	}

	waitCtx, cancel := context.WithTimeout(ctx, time.Minute*5)
	defer cancel()

	js, err := job.Wait(waitCtx)
	if err != nil {
		return errors.Annotate(err, "waiting for cluster summary analysis to complete").Err()
	}
	if js.Err() != nil {
		return errors.Annotate(err, "cluster summary analysis failed").Err()
	}
	return nil
}

// PurgeStaleRows purges stale clustered failure rows from the table.
// Stale rows are those rows which have been superseded by a new row with a later
// version, or where the latest version of the row has the row not included in a
// cluster.
// This is necessary for:
//   - Our QueryClusterSummaries query, which for performance reasons (UI-interactive)
//     does not do filtering to fetch the latest version of rows and instead uses all
//     rows.
//   - Keeping the size of the BigQuery table to a minimum.
//
// We currently only purge the last 7 days to keep purging costs to a minimum and
// as this is as far as QueryClusterSummaries looks back.
func (c *Client) PurgeStaleRows(ctx context.Context, luciProject string) error {
	datasetID, err := bqutil.DatasetForProject(luciProject)
	if err != nil {
		return errors.Annotate(err, "getting dataset").Err()
	}
	dataset := c.client.Dataset(datasetID)

	// If something goes wrong with this statement it deletes everything
	// for some reason, the system can be restored as follows:
	// - Fix the statement.
	// - Bump the algorithm version on all algorithms, to trigger a
	//   re-clustering and re-export of all test results.
	q := c.client.Query(`
		DELETE FROM clustered_failures cf1
		WHERE
			cf1.partition_time > TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 7 DAY) AND
			-- Not in the streaming buffer. Streaming buffer keeps up to
			-- 30 minutes of data. We use 40 minutes here to allow some
			-- margin as our last_updated timestamp is the timestamp
			-- the chunk was committed in Spanner and export to BigQuery
			-- can be delayed from that.
			cf1.last_updated < TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 40 MINUTE) AND
			(
				-- Not the latest (cluster, test result) entry.
				cf1.last_updated < (SELECT MAX(cf2.last_updated)
								FROM clustered_failures cf2
								WHERE cf2.partition_time > TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 7 DAY)
									AND cf2.partition_time = cf1.partition_time
									AND cf2.cluster_algorithm = cf1.cluster_algorithm
									AND cf2.cluster_id = cf1.cluster_id
									AND cf2.chunk_id = cf1.chunk_id
									AND cf2.chunk_index = cf1.chunk_index
									)
				-- Or is the latest (cluster, test result) entry, but test result
				-- is no longer in cluster.
				OR NOT COALESCE(cf1.is_included, FALSE)
			)
	`)
	q.DefaultDatasetID = dataset.DatasetID

	job, err := q.Run(ctx)
	if err != nil {
		return errors.Annotate(err, "purge stale rows").Err()
	}

	waitCtx, cancel := context.WithTimeout(ctx, time.Minute*5)
	defer cancel()

	js, err := job.Wait(waitCtx)
	if err != nil {
		// BigQuery specifies that rows are kept in the streaming buffer for
		// 30 minutes, but sometimes exceeds this SLO. We could be less
		// aggressive at deleting rows, but that would make the average-case
		// experience worse. These errors should only occur occasionally,
		// so it is better to ignore them.
		if strings.Contains(err.Error(), "would affect rows in the streaming buffer, which is not supported") {
			logging.Warningf(ctx, "Row purge failed for %v because rows were in the streaming buffer for over 30 minutes. "+
				"If this message occurs more than 25 percent of the time, it should be investigated.", luciProject)
			return nil
		}
		return errors.Annotate(err, "waiting for stale row purge to complete").Err()
	}
	if js.Err() != nil {
		return errors.Annotate(err, "purge stale rows failed").Err()
	}
	return nil
}

// EmptyCluster returns a Cluster entry for a cluster without any
// clustered failures.
func EmptyCluster(clusterID clustering.ClusterID) *Cluster {
	emptyCluster := &Cluster{
		ClusterID:    clusterID,
		MetricValues: make(map[metrics.ID]metrics.TimewiseCounts),
	}
	for _, metric := range metrics.ComputedMetrics {
		// Because there are no failures in the cluster, all metrics are zero-valued.
		emptyCluster.MetricValues[metric.ID] = metrics.TimewiseCounts{}
	}
	return emptyCluster
}

// ReadCluster reads information about a cluster.
// If the dataset for the LUCI project does not exist, returns ProjectNotExistsErr.
// If information for the cluster could not be found (e.g. because there are no examples),
// returns an empty cluster.
func (c *Client) ReadCluster(ctx context.Context, luciProject string, clusterID clustering.ClusterID) (cl *Cluster, err error) {
	_, s := trace.StartSpan(ctx, "go.chromium.org/luci/analysis/internal/analysis/ReadCluster")
	s.Attribute("project", luciProject)
	defer func() { s.End(err) }()

	whereClause := `cluster_algorithm = @clusterAlgorithm AND cluster_id = @clusterID`
	params := []bigquery.QueryParameter{
		{Name: "clusterAlgorithm", Value: clusterID.Algorithm},
		{Name: "clusterID", Value: clusterID.ID},
	}

	clusters, err := c.readClustersWhere(ctx, luciProject, whereClause, params)
	if err != nil {
		return nil, err
	}
	if len(clusters) == 0 {
		return EmptyCluster(clusterID), nil
	}
	return clusters[0], nil
}

// ImpactfulClusterReadOptions specifies options for ReadImpactfulClusters().
type ImpactfulClusterReadOptions struct {
	// Project is the LUCI Project for which analysis is being performed.
	Project string
	// Thresholds is the set of thresholds, which if any are met
	// or exceeded, should result in the cluster being returned.
	// Thresholds are applied based on the residual actual
	// cluster impact.
	Thresholds *configpb.ImpactThreshold
	// AlwaysIncludeBugClusters controls whether to include analysis for all
	// bug clusters.
	AlwaysIncludeBugClusters bool
}

// ReadImpactfulClusters reads clusters exceeding specified impact metrics, or are otherwise
// nominated to be read.
func (c *Client) ReadImpactfulClusters(ctx context.Context, opts ImpactfulClusterReadOptions) (cs []*Cluster, err error) {
	_, s := trace.StartSpan(ctx, "go.chromium.org/luci/analysis/internal/analysis/ReadImpactfulClusters")
	s.Attribute("project", opts.Project)
	defer func() { s.End(err) }()

	if opts.Thresholds == nil {
		return nil, errors.New("thresholds must be specified")
	}

	whereCriticalFailuresExonerated, cfeParams := whereThresholdsMet(metrics.CriticalFailuresExonerated, opts.Thresholds.CriticalFailuresExonerated)
	whereFailures, failuresParams := whereThresholdsMet(metrics.Failures, opts.Thresholds.TestResultsFailed)
	wherePresubmits, presubmitParams := whereThresholdsMet(metrics.HumanClsFailedPresubmit, opts.Thresholds.PresubmitRunsFailed)

	whereClause := `(` + whereCriticalFailuresExonerated + `) OR (` + whereFailures + `)
		    OR (` + wherePresubmits + `)
		    OR (@alwaysIncludeBugClusters AND cluster_algorithm = @ruleAlgorithmName)`
	params := []bigquery.QueryParameter{
		{
			Name:  "ruleAlgorithmName",
			Value: rulesalgorithm.AlgorithmName,
		},
		{
			Name:  "alwaysIncludeBugClusters",
			Value: opts.AlwaysIncludeBugClusters,
		},
	}
	params = append(params, cfeParams...)
	params = append(params, failuresParams...)
	params = append(params, presubmitParams...)
	return c.readClustersWhere(ctx, opts.Project, whereClause, params)
}

// metricExpression returns a SQL expression for the given metric
// in the cluster_summaries table. The type of the result will be NULLABLE INTEGER.
// The value will be NULL only if the metric is not in the underlying table.
func metricExpression(metric metrics.Definition, basis metrics.CalculationBasis) string {
	jsonColumnName := fmt.Sprintf("metrics_%s", basis.ColumnSuffix())
	return fmt.Sprintf("INT64(%s.%s)", jsonColumnName, metric.BaseColumnName)
}

func (c *Client) readClustersWhere(ctx context.Context, project, whereClause string, params []bigquery.QueryParameter) ([]*Cluster, error) {
	var selectList []string
	for _, metric := range metrics.ComputedMetrics {
		for _, calculationBasis := range metrics.CalculationBases {
			outputColumnName := metric.ColumnName(calculationBasis.ColumnSuffix())
			metricSelect := fmt.Sprintf("%s AS %s,", metricExpression(metric, calculationBasis), outputColumnName)
			selectList = append(selectList, metricSelect)
		}
	}

	dataset, err := bqutil.DatasetForProject(project)
	if err != nil {
		return nil, errors.Annotate(err, "getting dataset").Err()
	}

	q := c.client.Query(`
		SELECT
			cluster_algorithm,
			cluster_id,
			` + strings.Join(selectList, "\n") + `
			distinct_user_cls_with_failures_residual_7d,
			distinct_user_cls_with_failures_7d,
			postsubmit_builds_with_failures_residual_7d,
			postsubmit_builds_with_failures_7d,
			example_failure_reason.primary_error_message as example_failure_reason,
			top_test_ids,
			realms,
			ARRAY(
				SELECT AS STRUCT value, count
				FROM UNNEST(top_monorail_components)
				WHERE value IS NOT NULL
			) as top_monorail_components
		FROM cluster_summaries
		WHERE ` + whereClause,
	)
	q.DefaultDatasetID = dataset
	q.Parameters = params

	job, err := q.Run(ctx)
	if err != nil {
		return nil, errors.Annotate(err, "querying clusters").Err()
	}
	it, err := job.Read(ctx)
	if err != nil {
		return nil, handleJobReadError(err)
	}
	clusters := []*Cluster{}
	for {
		var rowVals rowLoader
		err := it.Next(&rowVals)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Annotate(err, "obtain next cluster row").Err()
		}
		row := &Cluster{}
		row.ClusterID = clustering.ClusterID{
			Algorithm: rowVals.String("cluster_algorithm"),
			ID:        rowVals.String("cluster_id"),
		}
		row.MetricValues = make(map[metrics.ID]metrics.TimewiseCounts)
		for _, metric := range metrics.ComputedMetrics {
			valid := true
			var timewiseCounts metrics.TimewiseCounts
			for _, calculationBasis := range metrics.CalculationBases {
				outputColumnName := metric.ColumnName(calculationBasis.ColumnSuffix())
				value := rowVals.NullInt64(outputColumnName)
				if !value.Valid {
					valid = false
					break
				}
				timewiseCounts.PutValue(value.Int64, calculationBasis)
			}
			if !valid {
				continue
			}
			row.MetricValues[metric.ID] = timewiseCounts
		}

		row.DistinctUserCLsWithFailures7d = metrics.Counts{
			Residual: rowVals.Int64("distinct_user_cls_with_failures_residual_7d"),
			Nominal:  rowVals.Int64("distinct_user_cls_with_failures_7d"),
		}
		row.PostsubmitBuildsWithFailures7d = metrics.Counts{
			Residual: rowVals.Int64("postsubmit_builds_with_failures_residual_7d"),
			Nominal:  rowVals.Int64("postsubmit_builds_with_failures_7d"),
		}
		row.ExampleFailureReason = rowVals.NullString("example_failure_reason")
		row.TopTestIDs = rowVals.TopCounts("top_test_ids")
		row.TopMonorailComponents = rowVals.TopCounts("top_monorail_components")
		row.Realms = rowVals.Strings("realms")

		if err := rowVals.Error(); err != nil {
			return nil, errors.Annotate(err, "marshalling cluster row").Err()
		}

		clusters = append(clusters, row)
	}
	return clusters, nil
}

func valueOrDefault(value *int64, defaultValue int64) int64 {
	if value != nil {
		return *value
	}
	return defaultValue
}

// whereThresholdsMet generates a SQL Where clause to query
// where a particular metric meets a given threshold.
func whereThresholdsMet(metric metrics.Definition, threshold *configpb.MetricThreshold) (string, []bigquery.QueryParameter) {
	if threshold == nil {
		threshold = &configpb.MetricThreshold{}
	}

	sql := fmt.Sprintf("%s >= @%s OR ", metricExpression(metric, metrics.OneDayResidualBasis), metric.ColumnName("1d")) +
		fmt.Sprintf("%s >= @%s OR ", metricExpression(metric, metrics.ThreeDayResidualBasis), metric.ColumnName("3d")) +
		fmt.Sprintf("%s >= @%s", metricExpression(metric, metrics.SevenDayResidualBasis), metric.ColumnName("7d"))
	parameters := []bigquery.QueryParameter{
		{
			Name:  metric.ColumnName("1d"),
			Value: valueOrDefault(threshold.OneDay, math.MaxInt64),
		},
		{
			Name:  metric.ColumnName("3d"),
			Value: valueOrDefault(threshold.ThreeDay, math.MaxInt64),
		},
		{
			Name:  metric.ColumnName("7d"),
			Value: valueOrDefault(threshold.SevenDay, math.MaxInt64),
		},
	}
	return sql, parameters
}

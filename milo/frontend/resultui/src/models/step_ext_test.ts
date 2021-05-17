// Copyright 2020 The LUCI Authors.
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

import { assert } from 'chai';
import { render } from 'lit-html';

import { renderMarkdown } from '../libs/markdown_utils';
import { BuildStatus } from '../services/buildbucket';
import { StepExt } from './step_ext';

describe('StepExt', () => {
  function createStep(name: string, status = BuildStatus.Success, summaryMarkdown = '') {
    return new StepExt({
      name,
      startTime: '2020-11-01T21:43:03.351951Z',
      status,
      summaryMarkdown,
    });
  }

  describe('should compute parent/child names and depth correctly', () => {
    it('for root step', () => {
      const step = createStep('child');
      assert.strictEqual(step.parentName, null);
      assert.strictEqual(step.selfName, 'child');
      assert.strictEqual(step.depth, 0);
    });

    it('for step with a parent', () => {
      const step = createStep('parent|child');
      assert.strictEqual(step.parentName, 'parent');
      assert.strictEqual(step.selfName, 'child');
      assert.strictEqual(step.depth, 1);
    });

    it('for step with a grandparent', () => {
      const step = createStep('grand-parent|parent|child');
      assert.strictEqual(step.parentName, 'grand-parent|parent');
      assert.strictEqual(step.selfName, 'child');
      assert.strictEqual(step.depth, 2);
    });
  });

  describe('succeededRecursively', () => {
    it('succeeded step with no children should return true', async () => {
      const step = createStep('child', BuildStatus.Success);
      assert.isTrue(step.succeededRecursively);
    });

    it('succeeded step with only succeeded children should return true', async () => {
      const step = createStep('child', BuildStatus.Success);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Success));
      assert.isTrue(step.succeededRecursively);
    });

    it('succeeded step with failed child should return false', async () => {
      const step = createStep('parent', BuildStatus.Success);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Failure));
      assert.isFalse(step.succeededRecursively);
    });

    it('succeeded step with failed child should return false', async () => {
      const step = createStep('parent', BuildStatus.Success);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Failure));
      assert.isFalse(step.succeededRecursively);
    });

    it('succeeded step with started child should return false', async () => {
      const step = createStep('parent', BuildStatus.Success);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Started));
      assert.isFalse(step.succeededRecursively);
    });

    it('succeeded step with scheduled child should return false', async () => {
      const step = createStep('parent', BuildStatus.Success);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Scheduled));
      assert.isFalse(step.succeededRecursively);
    });

    it('succeeded step with canceled child should return false', async () => {
      const step = createStep('parent', BuildStatus.Success);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Canceled));
      assert.isFalse(step.succeededRecursively);
    });

    it('failed step with no children should return false', async () => {
      const step = createStep('child', BuildStatus.Failure);
      assert.isFalse(step.succeededRecursively);
    });

    it('failed step with succeeded children should return false', async () => {
      const step = createStep('parent', BuildStatus.Failure);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Success));
      assert.isFalse(step.succeededRecursively);
    });

    it('failed step with failed children should return false', async () => {
      const step = createStep('parent', BuildStatus.Failure);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Failure));
      assert.isFalse(step.succeededRecursively);
    });
  });

  describe('failed', () => {
    it('succeeded step with no children should return false', async () => {
      const step = createStep('child', BuildStatus.Success);
      assert.isFalse(step.failed);
    });

    it('succeeded step with only succeeded children should return false', async () => {
      const step = createStep('child', BuildStatus.Success);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Success));
      assert.isFalse(step.failed);
    });

    it('succeeded step with failed child should return true', async () => {
      const step = createStep('parent', BuildStatus.Success);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Failure));
      assert.isTrue(step.failed);
    });

    it('failed step with no children should return true', async () => {
      const step = createStep('child', BuildStatus.Failure);
      assert.isTrue(step.failed);
    });

    it('infra-failed step with no children should return true', async () => {
      const step = createStep('child', BuildStatus.InfraFailure);
      assert.isTrue(step.failed);
    });

    it('canceled step with no children should return false', async () => {
      const step = createStep('child', BuildStatus.Canceled);
      assert.isFalse(step.failed);
    });

    it('scheduled step with no children should return false', async () => {
      const step = createStep('child', BuildStatus.Scheduled);
      assert.isFalse(step.failed);
    });

    it('started step with no children should return false', async () => {
      const step = createStep('child', BuildStatus.Started);
      assert.isFalse(step.failed);
    });

    it('failed step with succeeded children should return true', async () => {
      const step = createStep('parent', BuildStatus.Failure);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Success));
      assert.isTrue(step.failed);
    });

    it('failed step with failed children should return true', async () => {
      const step = createStep('parent', BuildStatus.Failure);
      step.children.push(createStep('parent|child1', BuildStatus.Success));
      step.children.push(createStep('parent|child2', BuildStatus.Failure));
      assert.isTrue(step.failed);
    });
  });

  describe('summary header and content should be split properly', () => {
    function getExpectedHeaderHTML(markdownBody: string): string {
      const container = document.createElement('div');
      // Wrap in a <p> and remove it later so <!----> are not injected.
      render(renderMarkdown(`<p>${markdownBody}</p>`), container);
      return container.firstElementChild!.innerHTML;
    }

    function getExpectedBodyHTML(markdownBody: string): string {
      const container = document.createElement('div');
      render(renderMarkdown(markdownBody), container);
      return container.innerHTML;
    }

    it('for no summary', async () => {
      const step = createStep('step', BuildStatus.Success, undefined);
      assert.strictEqual(step.header, null);
      assert.strictEqual(step.summary, null);
    });

    it('for empty summary', async () => {
      const step = createStep('step', BuildStatus.Success, '');
      assert.strictEqual(step.header, null);
      assert.strictEqual(step.summary, null);
    });

    it('for text summary', async () => {
      const step = createStep('step', BuildStatus.Success, 'this is some text');
      assert.strictEqual(step.header?.innerHTML, 'this is some text');
      assert.strictEqual(step.summary, null);
    });

    it('for header and content separated by <br/>', async () => {
      const step = createStep('step', BuildStatus.Success, 'header<br/>content');
      assert.strictEqual(step.header?.innerHTML, getExpectedHeaderHTML('header'));
      assert.strictEqual(step.summary?.innerHTML, getExpectedBodyHTML('content'));
    });

    it('for header and content separated by <br/>, header is empty', async () => {
      const step = createStep('step', BuildStatus.Success, '<br/>body');
      assert.strictEqual(step.header, null);
      assert.strictEqual(step.summary?.innerHTML, getExpectedBodyHTML('body'));
    });

    it('for header and content separated by <br/>, body is empty', async () => {
      const step = createStep('step', BuildStatus.Success, 'header<br/>');
      assert.strictEqual(step.header?.innerHTML, getExpectedHeaderHTML('header'));
      assert.strictEqual(step.summary, null);
    });

    it('for header and content separated by <br/>, header is a link', async () => {
      const step = createStep('step', BuildStatus.Success, '<a href="http://google.com">Link</a><br/>content');
      assert.strictEqual(step.header?.innerHTML, getExpectedHeaderHTML('<a href="http://google.com">Link</a>'));
      assert.strictEqual(step.summary?.innerHTML, getExpectedBodyHTML('content'));
    });

    it('for header and content separated by <br/>, header has some inline elements', async () => {
      const step = createStep(
        'step',
        BuildStatus.Success,
        '<span>span</span><i>i</i><b>b</b><strong>strong</strong><br/>content'
      );
      assert.strictEqual(
        step.header?.innerHTML,
        getExpectedHeaderHTML('<span>span</span><i>i</i><b>b</b><strong>strong</strong>')
      );
      assert.strictEqual(step.summary?.innerHTML, getExpectedBodyHTML('content'));
    });

    it('for header and content separated by <br/>, header is a list', async () => {
      const step = createStep('step', BuildStatus.Success, '<ul><li>item</li></ul><br/>content');
      assert.strictEqual(step.header, null);
      assert.strictEqual(step.summary?.innerHTML, getExpectedBodyHTML('<ul><li>item</li></ul><br/>content'));
    });

    it('for header is a list', async () => {
      const step = createStep('step', BuildStatus.Success, '<ul><li>item1</li><li>item2</li></ul>');
      assert.strictEqual(step.header, null);
      assert.strictEqual(step.summary?.innerHTML, getExpectedBodyHTML('<ul><li>item1</li><li>item2</li></ul>'));
    });

    it('for <br/> is contained in <div>', async () => {
      const step = createStep('step', BuildStatus.Success, '<div>header<br/>other</div>content');
      assert.strictEqual(step.header?.innerHTML, getExpectedHeaderHTML('header'));
      assert.strictEqual(step.summary?.innerHTML, getExpectedBodyHTML('<div>other</div>content'));
    });

    it('for <br/> is contained in some nested tags', async () => {
      const step = createStep('step', BuildStatus.Success, '<div><div>header<br/>other</div></div>content');
      assert.strictEqual(step.header, null);
      assert.strictEqual(step.summary?.innerHTML, getExpectedBodyHTML('<div><div>header<br/>other</div></div>content'));
    });
  });
});

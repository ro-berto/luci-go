// Copyright 2020 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/swarming/proto/config/realms.proto

package configpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Realm permissions used in Swarming.
// The enums are converted to string names using get_permission_name
// in server/realms.py
// NEXT_ID: 12
type RealmPermission int32

const (
	RealmPermission_REALM_PERMISSION_UNSPECIFIED RealmPermission = 0
	// Permission 'swarming.pools.createTask'
	// This is required to create a task in the pool.
	// It will be checked at the new task API.
	RealmPermission_REALM_PERMISSION_POOLS_CREATE_TASK RealmPermission = 1
	// Permission 'swarming.pools.listTasks'
	// This is required to list/count tasks in the pool.
	// It will be checked at tasks list/count APIs.
	RealmPermission_REALM_PERMISSION_POOLS_LIST_TASKS RealmPermission = 4
	// Permission 'swarming.pools.cancelTask'
	// This is required to cancel a task in the pool.
	// It will be checked at task cancel APIs.
	RealmPermission_REALM_PERMISSION_POOLS_CANCEL_TASK RealmPermission = 5
	// Permission 'swarming.pools.createBot'
	// This is required to create a bot in the pool.
	// It will be checked at bot bootstrap, bot code APIs.
	RealmPermission_REALM_PERMISSION_POOLS_CREATE_BOT RealmPermission = 6
	// Permission 'swarming.pools.listBots'
	// This is required to list/count bots in the pool.
	// It will be checked at bots list/count APIs.
	RealmPermission_REALM_PERMISSION_POOLS_LIST_BOTS RealmPermission = 7
	// Permission 'swarming.pools.terminateBot'
	// This is required to terminate a bot in the pool.
	// It will be checked at bot terminate API.
	RealmPermission_REALM_PERMISSION_POOLS_TERMINATE_BOT RealmPermission = 8
	// Permission 'swarming.pools.deleteBot'
	// This is required to delete a bot in the pool.
	// It will be checked at bot delete API.
	RealmPermission_REALM_PERMISSION_POOLS_DELETE_BOT RealmPermission = 9
	// Permission 'swarming.tasks.createInRealm'
	// This is required to create a task in the realm.
	// It will be checked at the new task API.
	RealmPermission_REALM_PERMISSION_TASKS_CREATE_IN_REALM RealmPermission = 2
	// Permission 'swarming.tasks.actAs'
	// This is required to use a task service account in the realm.
	// It will be checked at the new task API.
	RealmPermission_REALM_PERMISSION_TASKS_ACT_AS RealmPermission = 3
	// Permission 'swarming.tasks.get'
	// This is required to get task request, result, outputs.
	// It will be checked at task request, result, stdout GET APIs.
	RealmPermission_REALM_PERMISSION_TASKS_GET RealmPermission = 10
	// Permission 'swarming.tasks.cancel'
	// This is required to cancel a task.
	// It will be checked at task cancel API.
	RealmPermission_REALM_PERMISSION_TASKS_CANCEL RealmPermission = 11
)

// Enum value maps for RealmPermission.
var (
	RealmPermission_name = map[int32]string{
		0:  "REALM_PERMISSION_UNSPECIFIED",
		1:  "REALM_PERMISSION_POOLS_CREATE_TASK",
		4:  "REALM_PERMISSION_POOLS_LIST_TASKS",
		5:  "REALM_PERMISSION_POOLS_CANCEL_TASK",
		6:  "REALM_PERMISSION_POOLS_CREATE_BOT",
		7:  "REALM_PERMISSION_POOLS_LIST_BOTS",
		8:  "REALM_PERMISSION_POOLS_TERMINATE_BOT",
		9:  "REALM_PERMISSION_POOLS_DELETE_BOT",
		2:  "REALM_PERMISSION_TASKS_CREATE_IN_REALM",
		3:  "REALM_PERMISSION_TASKS_ACT_AS",
		10: "REALM_PERMISSION_TASKS_GET",
		11: "REALM_PERMISSION_TASKS_CANCEL",
	}
	RealmPermission_value = map[string]int32{
		"REALM_PERMISSION_UNSPECIFIED":           0,
		"REALM_PERMISSION_POOLS_CREATE_TASK":     1,
		"REALM_PERMISSION_POOLS_LIST_TASKS":      4,
		"REALM_PERMISSION_POOLS_CANCEL_TASK":     5,
		"REALM_PERMISSION_POOLS_CREATE_BOT":      6,
		"REALM_PERMISSION_POOLS_LIST_BOTS":       7,
		"REALM_PERMISSION_POOLS_TERMINATE_BOT":   8,
		"REALM_PERMISSION_POOLS_DELETE_BOT":      9,
		"REALM_PERMISSION_TASKS_CREATE_IN_REALM": 2,
		"REALM_PERMISSION_TASKS_ACT_AS":          3,
		"REALM_PERMISSION_TASKS_GET":             10,
		"REALM_PERMISSION_TASKS_CANCEL":          11,
	}
)

func (x RealmPermission) Enum() *RealmPermission {
	p := new(RealmPermission)
	*p = x
	return p
}

func (x RealmPermission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RealmPermission) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_swarming_proto_config_realms_proto_enumTypes[0].Descriptor()
}

func (RealmPermission) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_swarming_proto_config_realms_proto_enumTypes[0]
}

func (x RealmPermission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RealmPermission.Descriptor instead.
func (RealmPermission) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDescGZIP(), []int{0}
}

var File_go_chromium_org_luci_swarming_proto_config_realms_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0xda, 0x03, 0x0a, 0x0f, 0x52,
	0x65, 0x61, 0x6c, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x1c, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x26, 0x0a, 0x22, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45, 0x41, 0x4c,
	0x4d, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4f,
	0x4c, 0x53, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x53, 0x10, 0x04, 0x12,
	0x26, 0x0a, 0x22, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45, 0x41, 0x4c, 0x4d,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4f, 0x4c,
	0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4f, 0x54, 0x10, 0x06, 0x12, 0x24,
	0x0a, 0x20, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x53, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x42, 0x4f,
	0x54, 0x53, 0x10, 0x07, 0x12, 0x28, 0x0a, 0x24, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x53, 0x5f, 0x54,
	0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4f, 0x54, 0x10, 0x08, 0x12, 0x25,
	0x0a, 0x21, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x4f, 0x4c, 0x53, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f,
	0x42, 0x4f, 0x54, 0x10, 0x09, 0x12, 0x2a, 0x0a, 0x26, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x53, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x10,
	0x02, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x5f,
	0x41, 0x53, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x53, 0x5f, 0x47,
	0x45, 0x54, 0x10, 0x0a, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x53, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x0b, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDescData = file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDesc
)

func file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDescData)
	})
	return file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDescData
}

var file_go_chromium_org_luci_swarming_proto_config_realms_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_swarming_proto_config_realms_proto_goTypes = []interface{}{
	(RealmPermission)(0), // 0: swarming.config.RealmPermission
}
var file_go_chromium_org_luci_swarming_proto_config_realms_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_swarming_proto_config_realms_proto_init() }
func file_go_chromium_org_luci_swarming_proto_config_realms_proto_init() {
	if File_go_chromium_org_luci_swarming_proto_config_realms_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_swarming_proto_config_realms_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_swarming_proto_config_realms_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_swarming_proto_config_realms_proto_enumTypes,
	}.Build()
	File_go_chromium_org_luci_swarming_proto_config_realms_proto = out.File
	file_go_chromium_org_luci_swarming_proto_config_realms_proto_rawDesc = nil
	file_go_chromium_org_luci_swarming_proto_config_realms_proto_goTypes = nil
	file_go_chromium_org_luci_swarming_proto_config_realms_proto_depIdxs = nil
}

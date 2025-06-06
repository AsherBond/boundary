// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/storage/iam/store/v1/role_grant.proto

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoleGrant struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// role_id is the ID of the role this is a part of
	// @inject_tag: `gorm:"primary_key"`
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" gorm:"primary_key"`
	// raw_grant is the string grant value as provided by the user
	// @inject_tag: `gorm:"default:null"`
	RawGrant string `protobuf:"bytes,3,opt,name=raw_grant,json=rawGrant,proto3" json:"raw_grant,omitempty" gorm:"default:null"`
	// canonical_grant is the canonical string representation of the grant value.
	// We use this as the unique constraint.
	// @inject_tag: `gorm:"primary_key"`
	CanonicalGrant string `protobuf:"bytes,4,opt,name=canonical_grant,json=canonicalGrant,proto3" json:"canonical_grant,omitempty" gorm:"primary_key"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RoleGrant) Reset() {
	*x = RoleGrant{}
	mi := &file_controller_storage_iam_store_v1_role_grant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleGrant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrant) ProtoMessage() {}

func (x *RoleGrant) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_role_grant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrant.ProtoReflect.Descriptor instead.
func (*RoleGrant) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_role_grant_proto_rawDescGZIP(), []int{0}
}

func (x *RoleGrant) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *RoleGrant) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleGrant) GetRawGrant() string {
	if x != nil {
		return x.RawGrant
	}
	return ""
}

func (x *RoleGrant) GetCanonicalGrant() string {
	if x != nil {
		return x.CanonicalGrant
	}
	return ""
}

var File_controller_storage_iam_store_v1_role_grant_proto protoreflect.FileDescriptor

const file_controller_storage_iam_store_v1_role_grant_proto_rawDesc = "" +
	"\n" +
	"0controller/storage/iam/store/v1/role_grant.proto\x12\x1fcontroller.storage.iam.store.v1\x1a/controller/storage/timestamp/v1/timestamp.proto\"\xb7\x01\n" +
	"\tRoleGrant\x12K\n" +
	"\vcreate_time\x18\x01 \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\x12\x17\n" +
	"\arole_id\x18\x02 \x01(\tR\x06roleId\x12\x1b\n" +
	"\traw_grant\x18\x03 \x01(\tR\brawGrant\x12'\n" +
	"\x0fcanonical_grant\x18\x04 \x01(\tR\x0ecanonicalGrantB8Z6github.com/hashicorp/boundary/internal/iam/store;storeb\x06proto3"

var (
	file_controller_storage_iam_store_v1_role_grant_proto_rawDescOnce sync.Once
	file_controller_storage_iam_store_v1_role_grant_proto_rawDescData []byte
)

func file_controller_storage_iam_store_v1_role_grant_proto_rawDescGZIP() []byte {
	file_controller_storage_iam_store_v1_role_grant_proto_rawDescOnce.Do(func() {
		file_controller_storage_iam_store_v1_role_grant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_storage_iam_store_v1_role_grant_proto_rawDesc), len(file_controller_storage_iam_store_v1_role_grant_proto_rawDesc)))
	})
	return file_controller_storage_iam_store_v1_role_grant_proto_rawDescData
}

var file_controller_storage_iam_store_v1_role_grant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_storage_iam_store_v1_role_grant_proto_goTypes = []any{
	(*RoleGrant)(nil),           // 0: controller.storage.iam.store.v1.RoleGrant
	(*timestamp.Timestamp)(nil), // 1: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_iam_store_v1_role_grant_proto_depIdxs = []int32{
	1, // 0: controller.storage.iam.store.v1.RoleGrant.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_controller_storage_iam_store_v1_role_grant_proto_init() }
func file_controller_storage_iam_store_v1_role_grant_proto_init() {
	if File_controller_storage_iam_store_v1_role_grant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_storage_iam_store_v1_role_grant_proto_rawDesc), len(file_controller_storage_iam_store_v1_role_grant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_iam_store_v1_role_grant_proto_goTypes,
		DependencyIndexes: file_controller_storage_iam_store_v1_role_grant_proto_depIdxs,
		MessageInfos:      file_controller_storage_iam_store_v1_role_grant_proto_msgTypes,
	}.Build()
	File_controller_storage_iam_store_v1_role_grant_proto = out.File
	file_controller_storage_iam_store_v1_role_grant_proto_goTypes = nil
	file_controller_storage_iam_store_v1_role_grant_proto_depIdxs = nil
}

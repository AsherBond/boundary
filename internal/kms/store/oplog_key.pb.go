// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/storage/kms/store/v1/oplog_key.proto

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

type OplogKey struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// private_id is used to access the key via an API
	// @inject_tag: gorm:"primary_key"
	PrivateId string `protobuf:"bytes,10,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key"`
	// root key id for the key
	// @inject_tag: `gorm:"default:null"`
	RootKeyId string `protobuf:"bytes,20,opt,name=root_key_id,json=rootKeyId,proto3" json:"root_key_id,omitempty" gorm:"default:null"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime    *timestamp.Timestamp `protobuf:"bytes,30,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OplogKey) Reset() {
	*x = OplogKey{}
	mi := &file_controller_storage_kms_store_v1_oplog_key_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OplogKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OplogKey) ProtoMessage() {}

func (x *OplogKey) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_kms_store_v1_oplog_key_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OplogKey.ProtoReflect.Descriptor instead.
func (*OplogKey) Descriptor() ([]byte, []int) {
	return file_controller_storage_kms_store_v1_oplog_key_proto_rawDescGZIP(), []int{0}
}

func (x *OplogKey) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *OplogKey) GetRootKeyId() string {
	if x != nil {
		return x.RootKeyId
	}
	return ""
}

func (x *OplogKey) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type OplogKeyVersion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// private_id is used to access the key version via an API
	// @inject_tag: gorm:"primary_key"
	PrivateId string `protobuf:"bytes,10,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key"`
	// id for the key version
	// @inject_tag: `gorm:"default:null"`
	OplogKeyId string `protobuf:"bytes,20,opt,name=oplog_key_id,json=oplogKeyId,proto3" json:"oplog_key_id,omitempty" gorm:"default:null"`
	// root_key_version_id of the version of the root key data.
	// @inject_tag: `gorm:"default:null"`
	RootKeyVersionId string `protobuf:"bytes,30,opt,name=root_key_version_id,json=rootKeyVersionId,proto3" json:"root_key_version_id,omitempty" gorm:"default:null"`
	// plain-text of the key data.  we are NOT storing this plain-text key
	// in the db.
	// @inject_tag: `gorm:"-" wrapping:"pt,key_data"`
	Key []byte `protobuf:"bytes,40,opt,name=key,proto3" json:"key,omitempty" gorm:"-" wrapping:"pt,key_data"`
	// ciphertext key data stored in the database
	// @inject_tag: `gorm:"column:key;not_null" wrapping:"ct,key_data"`
	CtKey []byte `protobuf:"bytes,50,opt,name=ct_key,json=ctKey,proto3" json:"ct_key,omitempty" gorm:"column:key;not_null" wrapping:"ct,key_data"`
	// version of the key data.  This is not used for optimistic locking, since
	// key versions are immutable.  It's just the version of the key.
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,60,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime    *timestamp.Timestamp `protobuf:"bytes,70,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OplogKeyVersion) Reset() {
	*x = OplogKeyVersion{}
	mi := &file_controller_storage_kms_store_v1_oplog_key_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OplogKeyVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OplogKeyVersion) ProtoMessage() {}

func (x *OplogKeyVersion) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_kms_store_v1_oplog_key_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OplogKeyVersion.ProtoReflect.Descriptor instead.
func (*OplogKeyVersion) Descriptor() ([]byte, []int) {
	return file_controller_storage_kms_store_v1_oplog_key_proto_rawDescGZIP(), []int{1}
}

func (x *OplogKeyVersion) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *OplogKeyVersion) GetOplogKeyId() string {
	if x != nil {
		return x.OplogKeyId
	}
	return ""
}

func (x *OplogKeyVersion) GetRootKeyVersionId() string {
	if x != nil {
		return x.RootKeyVersionId
	}
	return ""
}

func (x *OplogKeyVersion) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *OplogKeyVersion) GetCtKey() []byte {
	if x != nil {
		return x.CtKey
	}
	return nil
}

func (x *OplogKeyVersion) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *OplogKeyVersion) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

var File_controller_storage_kms_store_v1_oplog_key_proto protoreflect.FileDescriptor

const file_controller_storage_kms_store_v1_oplog_key_proto_rawDesc = "" +
	"\n" +
	"/controller/storage/kms/store/v1/oplog_key.proto\x12\x1fcontroller.storage.kms.store.v1\x1a/controller/storage/timestamp/v1/timestamp.proto\"\x96\x01\n" +
	"\bOplogKey\x12\x1d\n" +
	"\n" +
	"private_id\x18\n" +
	" \x01(\tR\tprivateId\x12\x1e\n" +
	"\vroot_key_id\x18\x14 \x01(\tR\trootKeyId\x12K\n" +
	"\vcreate_time\x18\x1e \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\"\x91\x02\n" +
	"\x0fOplogKeyVersion\x12\x1d\n" +
	"\n" +
	"private_id\x18\n" +
	" \x01(\tR\tprivateId\x12 \n" +
	"\foplog_key_id\x18\x14 \x01(\tR\n" +
	"oplogKeyId\x12-\n" +
	"\x13root_key_version_id\x18\x1e \x01(\tR\x10rootKeyVersionId\x12\x10\n" +
	"\x03key\x18( \x01(\fR\x03key\x12\x15\n" +
	"\x06ct_key\x182 \x01(\fR\x05ctKey\x12\x18\n" +
	"\aversion\x18< \x01(\rR\aversion\x12K\n" +
	"\vcreate_time\x18F \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTimeB8Z6github.com/hashicorp/boundary/internal/kms/store;storeb\x06proto3"

var (
	file_controller_storage_kms_store_v1_oplog_key_proto_rawDescOnce sync.Once
	file_controller_storage_kms_store_v1_oplog_key_proto_rawDescData []byte
)

func file_controller_storage_kms_store_v1_oplog_key_proto_rawDescGZIP() []byte {
	file_controller_storage_kms_store_v1_oplog_key_proto_rawDescOnce.Do(func() {
		file_controller_storage_kms_store_v1_oplog_key_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_storage_kms_store_v1_oplog_key_proto_rawDesc), len(file_controller_storage_kms_store_v1_oplog_key_proto_rawDesc)))
	})
	return file_controller_storage_kms_store_v1_oplog_key_proto_rawDescData
}

var file_controller_storage_kms_store_v1_oplog_key_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_storage_kms_store_v1_oplog_key_proto_goTypes = []any{
	(*OplogKey)(nil),            // 0: controller.storage.kms.store.v1.OplogKey
	(*OplogKeyVersion)(nil),     // 1: controller.storage.kms.store.v1.OplogKeyVersion
	(*timestamp.Timestamp)(nil), // 2: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_kms_store_v1_oplog_key_proto_depIdxs = []int32{
	2, // 0: controller.storage.kms.store.v1.OplogKey.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 1: controller.storage.kms.store.v1.OplogKeyVersion.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_storage_kms_store_v1_oplog_key_proto_init() }
func file_controller_storage_kms_store_v1_oplog_key_proto_init() {
	if File_controller_storage_kms_store_v1_oplog_key_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_storage_kms_store_v1_oplog_key_proto_rawDesc), len(file_controller_storage_kms_store_v1_oplog_key_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_kms_store_v1_oplog_key_proto_goTypes,
		DependencyIndexes: file_controller_storage_kms_store_v1_oplog_key_proto_depIdxs,
		MessageInfos:      file_controller_storage_kms_store_v1_oplog_key_proto_msgTypes,
	}.Build()
	File_controller_storage_kms_store_v1_oplog_key_proto = out.File
	file_controller_storage_kms_store_v1_oplog_key_proto_goTypes = nil
	file_controller_storage_kms_store_v1_oplog_key_proto_depIdxs = nil
}

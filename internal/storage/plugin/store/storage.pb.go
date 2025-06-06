// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/storage/storage/plugin/store/v1/storage.proto

// Package store provides protobufs for storing types in the storage plugin
// package.

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
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

type StorageBucket struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// The scope_id of the owning scope and must be set.
	// @inject_tag: `gorm:"not_null"`
	ScopeId string `protobuf:"bytes,2,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"not_null"`
	// name is optional. If set, it must be globally unique.
	// @inject_tag: `gorm:"default:null"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" gorm:"default:null"`
	// description is optional.
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	// The plugin_id is the public id of the plugin managing this storage bucket.
	// @inject_tag: `gorm:"not_null"`
	PluginId string `protobuf:"bytes,8,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty" gorm:"not_null"`
	// The name of the external storage bucket (as opposed to the name within boundary)
	// @inject_tag: `gorm:"not_null"`
	BucketName string `protobuf:"bytes,9,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty" gorm:"not_null"`
	// @inject_tag: `gorm:"default:null"`
	BucketPrefix string `protobuf:"bytes,10,opt,name=bucket_prefix,json=bucketPrefix,proto3" json:"bucket_prefix,omitempty" gorm:"default:null"`
	// A boolean expression that allows filtering the workers that can handle a storage buckets
	// @inject_tag: `gorm:"not_null"`
	WorkerFilter string `protobuf:"bytes,11,opt,name=worker_filter,json=workerFilter,proto3" json:"worker_filter,omitempty" gorm:"not_null"`
	// attributes is a json formatted field.
	// @inject_tag: `gorm:"not_null"`
	Attributes []byte `protobuf:"bytes,12,opt,name=attributes,proto3" json:"attributes,omitempty" gorm:"not_null"`
	// secrets_hmac is a sha256-hmac of the unencrypted secrets stored in the db.
	// @inject_tag: `gorm:"default:null"`
	SecretsHmac []byte `protobuf:"bytes,13,opt,name=secrets_hmac,json=secretsHmac,proto3" json:"secrets_hmac,omitempty" gorm:"default:null"`
	// storage_bucket_credential_id is the private key of the storage bucket credential.
	// @inject_tag: `gorm:"default:null"`
	StorageBucketCredentialId string `protobuf:"bytes,14,opt,name=storage_bucket_credential_id,json=storageBucketCredentialId,proto3" json:"storage_bucket_credential_id,omitempty" gorm:"default:null"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *StorageBucket) Reset() {
	*x = StorageBucket{}
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StorageBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucket) ProtoMessage() {}

func (x *StorageBucket) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucket.ProtoReflect.Descriptor instead.
func (*StorageBucket) Descriptor() ([]byte, []int) {
	return file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescGZIP(), []int{0}
}

func (x *StorageBucket) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *StorageBucket) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *StorageBucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StorageBucket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *StorageBucket) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *StorageBucket) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *StorageBucket) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *StorageBucket) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *StorageBucket) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *StorageBucket) GetBucketPrefix() string {
	if x != nil {
		return x.BucketPrefix
	}
	return ""
}

func (x *StorageBucket) GetWorkerFilter() string {
	if x != nil {
		return x.WorkerFilter
	}
	return ""
}

func (x *StorageBucket) GetAttributes() []byte {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *StorageBucket) GetSecretsHmac() []byte {
	if x != nil {
		return x.SecretsHmac
	}
	return nil
}

func (x *StorageBucket) GetStorageBucketCredentialId() string {
	if x != nil {
		return x.StorageBucketCredentialId
	}
	return ""
}

type StorageBucketCredential struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// private_id is a surrogate key suitable for use via API.
	// @inject_tag: `gorm:"primary_key"`
	PrivateId string `protobuf:"bytes,1,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key"`
	// storage_bucket_id is the public id of the storage bucket.
	// @inject_tag: `gorm:"not_null"`
	StorageBucketId string `protobuf:"bytes,2,opt,name=storage_bucket_id,json=storageBucketId,proto3" json:"storage_bucket_id,omitempty" gorm:"not_null"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StorageBucketCredential) Reset() {
	*x = StorageBucketCredential{}
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StorageBucketCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketCredential) ProtoMessage() {}

func (x *StorageBucketCredential) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketCredential.ProtoReflect.Descriptor instead.
func (*StorageBucketCredential) Descriptor() ([]byte, []int) {
	return file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescGZIP(), []int{1}
}

func (x *StorageBucketCredential) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *StorageBucketCredential) GetStorageBucketId() string {
	if x != nil {
		return x.StorageBucketId
	}
	return ""
}

type StorageBucketCredentialEnvironmental struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// private_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PrivateId string `protobuf:"bytes,1,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key"`
	// storage_bucket_id is the public id of the storage bucket.
	// @inject_tag: `gorm:"not_null"`
	StorageBucketId string `protobuf:"bytes,2,opt,name=storage_bucket_id,json=storageBucketId,proto3" json:"storage_bucket_id,omitempty" gorm:"not_null"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StorageBucketCredentialEnvironmental) Reset() {
	*x = StorageBucketCredentialEnvironmental{}
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StorageBucketCredentialEnvironmental) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketCredentialEnvironmental) ProtoMessage() {}

func (x *StorageBucketCredentialEnvironmental) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketCredentialEnvironmental.ProtoReflect.Descriptor instead.
func (*StorageBucketCredentialEnvironmental) Descriptor() ([]byte, []int) {
	return file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescGZIP(), []int{2}
}

func (x *StorageBucketCredentialEnvironmental) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *StorageBucketCredentialEnvironmental) GetStorageBucketId() string {
	if x != nil {
		return x.StorageBucketId
	}
	return ""
}

type StorageBucketCredentialManagedSecret struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// private_id is a surrogate key suitable for use via API.
	// @inject_tag: `gorm:"primary_key"`
	PrivateId string `protobuf:"bytes,1,opt,name=private_id,json=privateId,proto3" json:"private_id,omitempty" gorm:"primary_key"`
	// storage_bucket_id is the public id of the storage bucket.
	// @inject_tag: `gorm:"not_null"`
	StorageBucketId string `protobuf:"bytes,2,opt,name=storage_bucket_id,json=storageBucketId,proto3" json:"storage_bucket_id,omitempty" gorm:"not_null"`
	// secrets is the plain-text of the secret data. We are not storing this plain-text
	// value in the database.
	// @inject_tag: `gorm:"-" wrapping:"pt,secrets_data"`
	Secrets []byte `protobuf:"bytes,3,opt,name=secrets,proto3" json:"secrets,omitempty" gorm:"-" wrapping:"pt,secrets_data"`
	// ct_secrets is the ciphertext of the secret data stored in the db.
	// @inject_tag: `gorm:"column:secrets_encrypted;not_null" wrapping:"ct,secrets_data"`
	CtSecrets []byte `protobuf:"bytes,4,opt,name=ct_secrets,json=ctSecrets,proto3" json:"ct_secrets,omitempty" gorm:"column:secrets_encrypted;not_null" wrapping:"ct,secrets_data"`
	// The key_id of the kms database key used for encrypting this entry.
	// It must be set.
	// @inject_tag: `gorm:"not_null"`
	KeyId         string `protobuf:"bytes,5,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty" gorm:"not_null"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StorageBucketCredentialManagedSecret) Reset() {
	*x = StorageBucketCredentialManagedSecret{}
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StorageBucketCredentialManagedSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBucketCredentialManagedSecret) ProtoMessage() {}

func (x *StorageBucketCredentialManagedSecret) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBucketCredentialManagedSecret.ProtoReflect.Descriptor instead.
func (*StorageBucketCredentialManagedSecret) Descriptor() ([]byte, []int) {
	return file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescGZIP(), []int{3}
}

func (x *StorageBucketCredentialManagedSecret) GetPrivateId() string {
	if x != nil {
		return x.PrivateId
	}
	return ""
}

func (x *StorageBucketCredentialManagedSecret) GetStorageBucketId() string {
	if x != nil {
		return x.StorageBucketId
	}
	return ""
}

func (x *StorageBucketCredentialManagedSecret) GetSecrets() []byte {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *StorageBucketCredentialManagedSecret) GetCtSecrets() []byte {
	if x != nil {
		return x.CtSecrets
	}
	return nil
}

func (x *StorageBucketCredentialManagedSecret) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

var File_controller_storage_storage_plugin_store_v1_storage_proto protoreflect.FileDescriptor

const file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc = "" +
	"\n" +
	"8controller/storage/storage/plugin/store/v1/storage.proto\x12*controller.storage.storage.plugin.store.v1\x1a*controller/custom_options/v1/options.proto\x1a/controller/storage/timestamp/v1/timestamp.proto\"\x92\x05\n" +
	"\rStorageBucket\x12\x1b\n" +
	"\tpublic_id\x18\x01 \x01(\tR\bpublicId\x12\x19\n" +
	"\bscope_id\x18\x02 \x01(\tR\ascopeId\x12$\n" +
	"\x04name\x18\x03 \x01(\tB\x10\xc2\xdd)\f\n" +
	"\x04Name\x12\x04nameR\x04name\x12@\n" +
	"\vdescription\x18\x04 \x01(\tB\x1e\xc2\xdd)\x1a\n" +
	"\vDescription\x12\vdescriptionR\vdescription\x12K\n" +
	"\vcreate_time\x18\x05 \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\x12K\n" +
	"\vupdate_time\x18\x06 \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"updateTime\x12\x18\n" +
	"\aversion\x18\a \x01(\rR\aversion\x12\x1b\n" +
	"\tplugin_id\x18\b \x01(\tR\bpluginId\x12\x1f\n" +
	"\vbucket_name\x18\t \x01(\tR\n" +
	"bucketName\x12#\n" +
	"\rbucket_prefix\x18\n" +
	" \x01(\tR\fbucketPrefix\x12F\n" +
	"\rworker_filter\x18\v \x01(\tB!\xc2\xdd)\x1d\n" +
	"\fWorkerFilter\x12\rworker_filterR\fworkerFilter\x12\x1e\n" +
	"\n" +
	"attributes\x18\f \x01(\fR\n" +
	"attributes\x12!\n" +
	"\fsecrets_hmac\x18\r \x01(\fR\vsecretsHmac\x12?\n" +
	"\x1cstorage_bucket_credential_id\x18\x0e \x01(\tR\x19storageBucketCredentialId\"d\n" +
	"\x17StorageBucketCredential\x12\x1d\n" +
	"\n" +
	"private_id\x18\x01 \x01(\tR\tprivateId\x12*\n" +
	"\x11storage_bucket_id\x18\x02 \x01(\tR\x0fstorageBucketId\"q\n" +
	"$StorageBucketCredentialEnvironmental\x12\x1d\n" +
	"\n" +
	"private_id\x18\x01 \x01(\tR\tprivateId\x12*\n" +
	"\x11storage_bucket_id\x18\x02 \x01(\tR\x0fstorageBucketId\"\xc1\x01\n" +
	"$StorageBucketCredentialManagedSecret\x12\x1d\n" +
	"\n" +
	"private_id\x18\x01 \x01(\tR\tprivateId\x12*\n" +
	"\x11storage_bucket_id\x18\x02 \x01(\tR\x0fstorageBucketId\x12\x18\n" +
	"\asecrets\x18\x03 \x01(\fR\asecrets\x12\x1d\n" +
	"\n" +
	"ct_secrets\x18\x04 \x01(\fR\tctSecrets\x12\x15\n" +
	"\x06key_id\x18\x05 \x01(\tR\x05keyIdBCZAgithub.com/hashicorp/boundary/internal/storage/plugin/store;storeb\x06proto3"

var (
	file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescOnce sync.Once
	file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescData []byte
)

func file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescGZIP() []byte {
	file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescOnce.Do(func() {
		file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc), len(file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc)))
	})
	return file_controller_storage_storage_plugin_store_v1_storage_proto_rawDescData
}

var file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controller_storage_storage_plugin_store_v1_storage_proto_goTypes = []any{
	(*StorageBucket)(nil),                        // 0: controller.storage.storage.plugin.store.v1.StorageBucket
	(*StorageBucketCredential)(nil),              // 1: controller.storage.storage.plugin.store.v1.StorageBucketCredential
	(*StorageBucketCredentialEnvironmental)(nil), // 2: controller.storage.storage.plugin.store.v1.StorageBucketCredentialEnvironmental
	(*StorageBucketCredentialManagedSecret)(nil), // 3: controller.storage.storage.plugin.store.v1.StorageBucketCredentialManagedSecret
	(*timestamp.Timestamp)(nil),                  // 4: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_storage_plugin_store_v1_storage_proto_depIdxs = []int32{
	4, // 0: controller.storage.storage.plugin.store.v1.StorageBucket.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	4, // 1: controller.storage.storage.plugin.store.v1.StorageBucket.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_storage_storage_plugin_store_v1_storage_proto_init() }
func file_controller_storage_storage_plugin_store_v1_storage_proto_init() {
	if File_controller_storage_storage_plugin_store_v1_storage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc), len(file_controller_storage_storage_plugin_store_v1_storage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_storage_plugin_store_v1_storage_proto_goTypes,
		DependencyIndexes: file_controller_storage_storage_plugin_store_v1_storage_proto_depIdxs,
		MessageInfos:      file_controller_storage_storage_plugin_store_v1_storage_proto_msgTypes,
	}.Build()
	File_controller_storage_storage_plugin_store_v1_storage_proto = out.File
	file_controller_storage_storage_plugin_store_v1_storage_proto_goTypes = nil
	file_controller_storage_storage_plugin_store_v1_storage_proto_depIdxs = nil
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: controller/storage/servers/store/v1/worker_auth.proto

// Package store provides protobufs for storing types in the pki package.

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
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

// WorkerAuth contains all fields related to an authorized Worker resource
type WorkerAuth struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The key id for this worker auth record, generated from the signing pub key
	// @inject_tag: `gorm:"primary_key"`
	WorkerKeyIdentifier string `protobuf:"bytes,10,opt,name=worker_key_identifier,json=workerKeyIdentifier,proto3" json:"worker_key_identifier,omitempty" gorm:"primary_key"`
	// The worker id this worker authentication record is for
	// @inject_tag: `gorm:"not_null"`
	WorkerId string `protobuf:"bytes,20,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty" gorm:"not_null"`
	// The worker's public signing key
	// @inject_tag: `gorm:"not_null"`
	WorkerSigningPubKey []byte `protobuf:"bytes,30,opt,name=worker_signing_pub_key,json=workerSigningPubKey,proto3" json:"worker_signing_pub_key,omitempty" gorm:"not_null"`
	// The worker's public encryption key
	// @inject_tag: `gorm:"not_null"`
	WorkerEncryptionPubKey []byte `protobuf:"bytes,40,opt,name=worker_encryption_pub_key,json=workerEncryptionPubKey,proto3" json:"worker_encryption_pub_key,omitempty" gorm:"not_null"`
	// The plain-text private key associated with this certificate. We are not storing this
	// in the database.
	// @inject_tag: gorm:"-" wrapping:"pt,private_key"
	ControllerEncryptionPrivKey []byte `protobuf:"bytes,50,opt,name=controller_encryption_priv_key,json=controllerEncryptionPrivKey,proto3" json:"controller_encryption_priv_key,omitempty" gorm:"-" wrapping:"pt,private_key"`
	// The private key associated with this certificate
	// This is a ciphertext field
	// @inject_tag: gorm:"column:controller_encryption_priv_key;not_null" wrapping:"ct,private_key"
	CtControllerEncryptionPrivKey []byte `protobuf:"bytes,51,opt,name=ct_controller_encryption_priv_key,json=ctControllerEncryptionPrivKey,proto3" json:"ct_controller_encryption_priv_key,omitempty" gorm:"column:controller_encryption_priv_key;not_null" wrapping:"ct,private_key"`
	// The id of the kms database key used for encrypting this entry.
	// @inject_tag: `gorm:"not_null"`
	KeyId string `protobuf:"bytes,60,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty" gorm:"not_null"`
	// Nonce used by a worker in authenticating
	// @inject_tag: `gorm:"default:null"`
	Nonce []byte `protobuf:"bytes,70,opt,name=nonce,proto3" json:"nonce,omitempty" gorm:"default:null"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,80,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,90,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// State of the worker auth record.
	// The only valid value is either current or previous
	// @inject_tag: `gorm:"not_null"`
	State         string `protobuf:"bytes,100,opt,name=state,proto3" json:"state,omitempty" gorm:"not_null"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkerAuth) Reset() {
	*x = WorkerAuth{}
	mi := &file_controller_storage_servers_store_v1_worker_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkerAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerAuth) ProtoMessage() {}

func (x *WorkerAuth) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_servers_store_v1_worker_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerAuth.ProtoReflect.Descriptor instead.
func (*WorkerAuth) Descriptor() ([]byte, []int) {
	return file_controller_storage_servers_store_v1_worker_auth_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerAuth) GetWorkerKeyIdentifier() string {
	if x != nil {
		return x.WorkerKeyIdentifier
	}
	return ""
}

func (x *WorkerAuth) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *WorkerAuth) GetWorkerSigningPubKey() []byte {
	if x != nil {
		return x.WorkerSigningPubKey
	}
	return nil
}

func (x *WorkerAuth) GetWorkerEncryptionPubKey() []byte {
	if x != nil {
		return x.WorkerEncryptionPubKey
	}
	return nil
}

func (x *WorkerAuth) GetControllerEncryptionPrivKey() []byte {
	if x != nil {
		return x.ControllerEncryptionPrivKey
	}
	return nil
}

func (x *WorkerAuth) GetCtControllerEncryptionPrivKey() []byte {
	if x != nil {
		return x.CtControllerEncryptionPrivKey
	}
	return nil
}

func (x *WorkerAuth) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *WorkerAuth) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *WorkerAuth) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *WorkerAuth) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *WorkerAuth) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

// WorkerCertBundle contains all fields related to a WorkerCertBundle resource
type WorkerCertBundle struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The public key of the issuing root certificate
	// @inject_tag: `gorm:"primary_key"`
	RootCertificatePublicKey []byte `protobuf:"bytes,10,opt,name=root_certificate_public_key,json=rootCertificatePublicKey,proto3" json:"root_certificate_public_key,omitempty" gorm:"primary_key"`
	// The WorkerAuth worker_key_identifier this cert bundle record is for
	// @inject_tag: `gorm:"primary_key"`
	WorkerKeyIdentifier string `protobuf:"bytes,20,opt,name=worker_key_identifier,json=workerKeyIdentifier,proto3" json:"worker_key_identifier,omitempty" gorm:"primary_key"`
	// CertBundle is the marshaled protobuf certificate bundle for a WorkerAuth
	// @inject_tag: `gorm:"not_null"`
	CertBundle    []byte `protobuf:"bytes,30,opt,name=cert_bundle,json=certBundle,proto3" json:"cert_bundle,omitempty" gorm:"not_null"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkerCertBundle) Reset() {
	*x = WorkerCertBundle{}
	mi := &file_controller_storage_servers_store_v1_worker_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkerCertBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerCertBundle) ProtoMessage() {}

func (x *WorkerCertBundle) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_servers_store_v1_worker_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerCertBundle.ProtoReflect.Descriptor instead.
func (*WorkerCertBundle) Descriptor() ([]byte, []int) {
	return file_controller_storage_servers_store_v1_worker_auth_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerCertBundle) GetRootCertificatePublicKey() []byte {
	if x != nil {
		return x.RootCertificatePublicKey
	}
	return nil
}

func (x *WorkerCertBundle) GetWorkerKeyIdentifier() string {
	if x != nil {
		return x.WorkerKeyIdentifier
	}
	return ""
}

func (x *WorkerCertBundle) GetCertBundle() []byte {
	if x != nil {
		return x.CertBundle
	}
	return nil
}

// WorkerAuthServerLedActivationToken contains all fields related to a
// WorkerAuthServerLedActivationToken resource
type WorkerAuthServerLedActivationToken struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The worker_id of the worker that this activates
	// @inject_tag: `gorm:"primary_key"`
	WorkerId string `protobuf:"bytes,10,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty" gorm:"primary_key"`
	// The token identifier, which is used for lookup
	// @inject_tag: `gorm:"not_null"`
	TokenId string `protobuf:"bytes,15,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty" gorm:"not_null"`
	// The creation time, encrypted to prevent tampering, as the time plus
	// existence of the record allows authorization
	// @inject_tag: `gorm:"not_null" wrapping:"ct,creation_time_data"`
	CreationTimeEncrypted []byte `protobuf:"bytes,20,opt,name=creation_time_encrypted,json=creationTimeEncrypted,proto3" json:"creation_time_encrypted,omitempty" gorm:"not_null" wrapping:"ct,creation_time_data"`
	// The plaintext bytes of the creation time, which are never stored. This is a
	// marshaled timestamppb.Timestamp.
	// @inject_tag: `gorm:"-" wrapping:"pt,creation_time_data"`
	CreationTime []byte `protobuf:"bytes,21,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty" gorm:"-" wrapping:"pt,creation_time_data"`
	// The key ID of the encrypting key
	// @inject_tag: `gorm:"not_null"`
	KeyId         string `protobuf:"bytes,40,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty" gorm:"not_null"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkerAuthServerLedActivationToken) Reset() {
	*x = WorkerAuthServerLedActivationToken{}
	mi := &file_controller_storage_servers_store_v1_worker_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkerAuthServerLedActivationToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerAuthServerLedActivationToken) ProtoMessage() {}

func (x *WorkerAuthServerLedActivationToken) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_servers_store_v1_worker_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerAuthServerLedActivationToken.ProtoReflect.Descriptor instead.
func (*WorkerAuthServerLedActivationToken) Descriptor() ([]byte, []int) {
	return file_controller_storage_servers_store_v1_worker_auth_proto_rawDescGZIP(), []int{2}
}

func (x *WorkerAuthServerLedActivationToken) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *WorkerAuthServerLedActivationToken) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *WorkerAuthServerLedActivationToken) GetCreationTimeEncrypted() []byte {
	if x != nil {
		return x.CreationTimeEncrypted
	}
	return nil
}

func (x *WorkerAuthServerLedActivationToken) GetCreationTime() []byte {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *WorkerAuthServerLedActivationToken) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

var File_controller_storage_servers_store_v1_worker_auth_proto protoreflect.FileDescriptor

var file_controller_storage_servers_store_v1_worker_auth_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x04,
	0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x32, 0x0a, 0x15,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x0a,
	0x16, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x12, 0x39, 0x0a, 0x19, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x16, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x43, 0x0a,
	0x1e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x76, 0x4b,
	0x65, 0x79, 0x12, 0x48, 0x0a, 0x21, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x69, 0x76, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1d, 0x63,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06,
	0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65,
	0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x46, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x10, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x3d,
	0x0a, 0x1b, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x18, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a,
	0x15, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x65, 0x72, 0x74, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x22, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x64, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6b, 0x65, 0x79, 0x49, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_servers_store_v1_worker_auth_proto_rawDescOnce sync.Once
	file_controller_storage_servers_store_v1_worker_auth_proto_rawDescData = file_controller_storage_servers_store_v1_worker_auth_proto_rawDesc
)

func file_controller_storage_servers_store_v1_worker_auth_proto_rawDescGZIP() []byte {
	file_controller_storage_servers_store_v1_worker_auth_proto_rawDescOnce.Do(func() {
		file_controller_storage_servers_store_v1_worker_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_servers_store_v1_worker_auth_proto_rawDescData)
	})
	return file_controller_storage_servers_store_v1_worker_auth_proto_rawDescData
}

var file_controller_storage_servers_store_v1_worker_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_controller_storage_servers_store_v1_worker_auth_proto_goTypes = []any{
	(*WorkerAuth)(nil),                         // 0: controller.storage.servers.store.v1.WorkerAuth
	(*WorkerCertBundle)(nil),                   // 1: controller.storage.servers.store.v1.WorkerCertBundle
	(*WorkerAuthServerLedActivationToken)(nil), // 2: controller.storage.servers.store.v1.WorkerAuthServerLedActivationToken
	(*timestamp.Timestamp)(nil),                // 3: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_servers_store_v1_worker_auth_proto_depIdxs = []int32{
	3, // 0: controller.storage.servers.store.v1.WorkerAuth.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	3, // 1: controller.storage.servers.store.v1.WorkerAuth.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_storage_servers_store_v1_worker_auth_proto_init() }
func file_controller_storage_servers_store_v1_worker_auth_proto_init() {
	if File_controller_storage_servers_store_v1_worker_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_storage_servers_store_v1_worker_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_servers_store_v1_worker_auth_proto_goTypes,
		DependencyIndexes: file_controller_storage_servers_store_v1_worker_auth_proto_depIdxs,
		MessageInfos:      file_controller_storage_servers_store_v1_worker_auth_proto_msgTypes,
	}.Build()
	File_controller_storage_servers_store_v1_worker_auth_proto = out.File
	file_controller_storage_servers_store_v1_worker_auth_proto_rawDesc = nil
	file_controller_storage_servers_store_v1_worker_auth_proto_goTypes = nil
	file_controller_storage_servers_store_v1_worker_auth_proto_depIdxs = nil
}

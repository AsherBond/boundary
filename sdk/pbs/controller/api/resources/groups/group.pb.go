// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/api/resources/groups/v1/group.proto

package groups

import (
	scopes "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type Member struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The ID of the member.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The Scope ID of the member.
	ScopeId       string `protobuf:"bytes,20,opt,name=scope_id,proto3" json:"scope_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Member) Reset() {
	*x = Member{}
	mi := &file_controller_api_resources_groups_v1_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_groups_v1_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_groups_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Member) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

// Group contains all fields related to a Group resource
type Group struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The ID of the Group.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// The ID of the scope of which this Group is a part.
	ScopeId string `protobuf:"bytes,20,opt,name=scope_id,proto3" json:"scope_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. Scope information for this Group.
	Scope *scopes.ScopeInfo `protobuf:"bytes,30,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional name for identification purposes.
	Name *wrapperspb.StringValue `protobuf:"bytes,40,opt,name=name,proto3" json:"name,omitempty" class:"public"` // @gotags: `class:"public"`
	// Optional user-set descripton for identification purposes.
	Description *wrapperspb.StringValue `protobuf:"bytes,50,opt,name=description,proto3" json:"description,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The time this resource was created.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,60,opt,name=created_time,proto3" json:"created_time,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The time this resource was last updated.
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,70,opt,name=updated_time,proto3" json:"updated_time,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Version is used in mutation requests, after the initial creation, to ensure this resource has not changed.
	// The mutation will fail if the version does not match the latest known good version.
	Version uint32 `protobuf:"varint,80,opt,name=version,proto3" json:"version,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. Contains the list of member IDs in this Group.
	MemberIds []string `protobuf:"bytes,90,rep,name=member_ids,proto3" json:"member_ids,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The members of this Group.
	Members []*Member `protobuf:"bytes,100,rep,name=members,proto3" json:"members,omitempty"`
	// Output only. The available actions on this resource for this user.
	AuthorizedActions []string `protobuf:"bytes,300,rep,name=authorized_actions,proto3" json:"authorized_actions,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_controller_api_resources_groups_v1_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_groups_v1_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_groups_v1_group_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *Group) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Group) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Group) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Group) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Group) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Group) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Group) GetMemberIds() []string {
	if x != nil {
		return x.MemberIds
	}
	return nil
}

func (x *Group) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Group) GetAuthorizedActions() []string {
	if x != nil {
		return x.AuthorizedActions
	}
	return nil
}

var File_controller_api_resources_groups_v1_group_proto protoreflect.FileDescriptor

const file_controller_api_resources_groups_v1_group_proto_rawDesc = "" +
	"\n" +
	".controller/api/resources/groups/v1/group.proto\x12\"controller.api.resources.groups.v1\x1a.controller/api/resources/scopes/v1/scope.proto\x1a*controller/custom_options/v1/options.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"4\n" +
	"\x06Member\x12\x0e\n" +
	"\x02id\x18\n" +
	" \x01(\tR\x02id\x12\x1a\n" +
	"\bscope_id\x18\x14 \x01(\tR\bscope_id\"\xd5\x04\n" +
	"\x05Group\x12\x0e\n" +
	"\x02id\x18\n" +
	" \x01(\tR\x02id\x12\x1a\n" +
	"\bscope_id\x18\x14 \x01(\tR\bscope_id\x12C\n" +
	"\x05scope\x18\x1e \x01(\v2-.controller.api.resources.scopes.v1.ScopeInfoR\x05scope\x12F\n" +
	"\x04name\x18( \x01(\v2\x1c.google.protobuf.StringValueB\x14\xa0\xda)\x01\xc2\xdd)\f\n" +
	"\x04name\x12\x04nameR\x04name\x12b\n" +
	"\vdescription\x182 \x01(\v2\x1c.google.protobuf.StringValueB\"\xa0\xda)\x01\xc2\xdd)\x1a\n" +
	"\vdescription\x12\vdescriptionR\vdescription\x12>\n" +
	"\fcreated_time\x18< \x01(\v2\x1a.google.protobuf.TimestampR\fcreated_time\x12>\n" +
	"\fupdated_time\x18F \x01(\v2\x1a.google.protobuf.TimestampR\fupdated_time\x12\x18\n" +
	"\aversion\x18P \x01(\rR\aversion\x12\x1e\n" +
	"\n" +
	"member_ids\x18Z \x03(\tR\n" +
	"member_ids\x12D\n" +
	"\amembers\x18d \x03(\v2*.controller.api.resources.groups.v1.MemberR\amembers\x12/\n" +
	"\x12authorized_actions\x18\xac\x02 \x03(\tR\x12authorized_actionsBNZLgithub.com/hashicorp/boundary/sdk/pbs/controller/api/resources/groups;groupsb\x06proto3"

var (
	file_controller_api_resources_groups_v1_group_proto_rawDescOnce sync.Once
	file_controller_api_resources_groups_v1_group_proto_rawDescData []byte
)

func file_controller_api_resources_groups_v1_group_proto_rawDescGZIP() []byte {
	file_controller_api_resources_groups_v1_group_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_groups_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_api_resources_groups_v1_group_proto_rawDesc), len(file_controller_api_resources_groups_v1_group_proto_rawDesc)))
	})
	return file_controller_api_resources_groups_v1_group_proto_rawDescData
}

var file_controller_api_resources_groups_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_api_resources_groups_v1_group_proto_goTypes = []any{
	(*Member)(nil),                 // 0: controller.api.resources.groups.v1.Member
	(*Group)(nil),                  // 1: controller.api.resources.groups.v1.Group
	(*scopes.ScopeInfo)(nil),       // 2: controller.api.resources.scopes.v1.ScopeInfo
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_controller_api_resources_groups_v1_group_proto_depIdxs = []int32{
	2, // 0: controller.api.resources.groups.v1.Group.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	3, // 1: controller.api.resources.groups.v1.Group.name:type_name -> google.protobuf.StringValue
	3, // 2: controller.api.resources.groups.v1.Group.description:type_name -> google.protobuf.StringValue
	4, // 3: controller.api.resources.groups.v1.Group.created_time:type_name -> google.protobuf.Timestamp
	4, // 4: controller.api.resources.groups.v1.Group.updated_time:type_name -> google.protobuf.Timestamp
	0, // 5: controller.api.resources.groups.v1.Group.members:type_name -> controller.api.resources.groups.v1.Member
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_controller_api_resources_groups_v1_group_proto_init() }
func file_controller_api_resources_groups_v1_group_proto_init() {
	if File_controller_api_resources_groups_v1_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_api_resources_groups_v1_group_proto_rawDesc), len(file_controller_api_resources_groups_v1_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_groups_v1_group_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_groups_v1_group_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_groups_v1_group_proto_msgTypes,
	}.Build()
	File_controller_api_resources_groups_v1_group_proto = out.File
	file_controller_api_resources_groups_v1_group_proto_goTypes = nil
	file_controller_api_resources_groups_v1_group_proto_depIdxs = nil
}

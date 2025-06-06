// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: controller/api/services/v1/user_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_GetUser_FullMethodName               = "/controller.api.services.v1.UserService/GetUser"
	UserService_ListUsers_FullMethodName             = "/controller.api.services.v1.UserService/ListUsers"
	UserService_CreateUser_FullMethodName            = "/controller.api.services.v1.UserService/CreateUser"
	UserService_UpdateUser_FullMethodName            = "/controller.api.services.v1.UserService/UpdateUser"
	UserService_DeleteUser_FullMethodName            = "/controller.api.services.v1.UserService/DeleteUser"
	UserService_AddUserAccounts_FullMethodName       = "/controller.api.services.v1.UserService/AddUserAccounts"
	UserService_SetUserAccounts_FullMethodName       = "/controller.api.services.v1.UserService/SetUserAccounts"
	UserService_RemoveUserAccounts_FullMethodName    = "/controller.api.services.v1.UserService/RemoveUserAccounts"
	UserService_ListResolvableAliases_FullMethodName = "/controller.api.services.v1.UserService/ListResolvableAliases"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// GetUser returns a stored User if present.  The provided request
	// must include the User ID for the User being retrieved. If
	// that ID is missing, malformed or reference a non existing
	// resource an error is returned.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// ListUsers returns a list of stored Users which exist inside the provided
	// scope. The request must include the scope ID for the Users being listed.
	// If the scope ID is missing, malformed, or reference a non existing scope,
	// an error is returned.
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// CreateUser creates and stores a User in boundary.  The provided
	// request must include the Scope id in which the User will be created.
	// If the Scope id is missing, malformed or references a non existing
	// resource, an error is returned.  If a name is provided that is in
	// use in another User in the same scope, an error is returned.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// UpdateUser updates an existing User in boundary.  The provided
	// User must not have any read only fields set.  The update mask must be
	// included in the request and contain at least 1 mutable field.  To unset
	// a field's value, include the field in the update mask and don't set it
	// in the provided User. An error is returned if either the User id is
	// missing or reference a non existing resource.  An error is also returned
	// if the request attempts to update the name to one that is already in use
	// in this Scope.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	// DeleteUser removes a User from Boundary. If the provided User ID
	// is malformed or not provided an error is returned.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	// AddUserAccounts adds Users as members to a group.  The provided request
	// must include the User id which the Account will be added to. If that id is
	// missing, malformed or references a non existing resource, an error is
	// returned. If any of the Accounts are associated with another User an
	// error is returned.
	AddUserAccounts(ctx context.Context, in *AddUserAccountsRequest, opts ...grpc.CallOption) (*AddUserAccountsResponse, error)
	// SetUserAccounts sets the Accounts associated with this User.
	// Any existing Accounts are removed if they are not included in this request.
	// The provided request must include the User ID which the Accounts will be
	// associated with.  Any Accounts not included in this request but previously
	// associated with this user will be disassociated.
	// If the User ID is missing, malformed or references a non existing resource,
	// an error is returned.
	// If any of the Accounts are associated with another User an error is returned.
	SetUserAccounts(ctx context.Context, in *SetUserAccountsRequest, opts ...grpc.CallOption) (*SetUserAccountsResponse, error)
	// RemoveUserAccounts removes Accounts from the specified User.
	// The provided request must include the User id which the Accounts
	// will be removed from. If the provided Account ids is not associated with the
	// provided User, an error is returned.
	RemoveUserAccounts(ctx context.Context, in *RemoveUserAccountsRequest, opts ...grpc.CallOption) (*RemoveUserAccountsResponse, error)
	// ListResolvableAliases returns a list of Aliases which point to a resource
	// for which the provided user id has some permission.
	// If missing or malformed an error is returned.
	ListResolvableAliases(ctx context.Context, in *ListResolvableAliasesRequest, opts ...grpc.CallOption) (*ListResolvableAliasesResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, UserService_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddUserAccounts(ctx context.Context, in *AddUserAccountsRequest, opts ...grpc.CallOption) (*AddUserAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserAccountsResponse)
	err := c.cc.Invoke(ctx, UserService_AddUserAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetUserAccounts(ctx context.Context, in *SetUserAccountsRequest, opts ...grpc.CallOption) (*SetUserAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetUserAccountsResponse)
	err := c.cc.Invoke(ctx, UserService_SetUserAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveUserAccounts(ctx context.Context, in *RemoveUserAccountsRequest, opts ...grpc.CallOption) (*RemoveUserAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveUserAccountsResponse)
	err := c.cc.Invoke(ctx, UserService_RemoveUserAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListResolvableAliases(ctx context.Context, in *ListResolvableAliasesRequest, opts ...grpc.CallOption) (*ListResolvableAliasesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResolvableAliasesResponse)
	err := c.cc.Invoke(ctx, UserService_ListResolvableAliases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	// GetUser returns a stored User if present.  The provided request
	// must include the User ID for the User being retrieved. If
	// that ID is missing, malformed or reference a non existing
	// resource an error is returned.
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// ListUsers returns a list of stored Users which exist inside the provided
	// scope. The request must include the scope ID for the Users being listed.
	// If the scope ID is missing, malformed, or reference a non existing scope,
	// an error is returned.
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// CreateUser creates and stores a User in boundary.  The provided
	// request must include the Scope id in which the User will be created.
	// If the Scope id is missing, malformed or references a non existing
	// resource, an error is returned.  If a name is provided that is in
	// use in another User in the same scope, an error is returned.
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// UpdateUser updates an existing User in boundary.  The provided
	// User must not have any read only fields set.  The update mask must be
	// included in the request and contain at least 1 mutable field.  To unset
	// a field's value, include the field in the update mask and don't set it
	// in the provided User. An error is returned if either the User id is
	// missing or reference a non existing resource.  An error is also returned
	// if the request attempts to update the name to one that is already in use
	// in this Scope.
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	// DeleteUser removes a User from Boundary. If the provided User ID
	// is malformed or not provided an error is returned.
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	// AddUserAccounts adds Users as members to a group.  The provided request
	// must include the User id which the Account will be added to. If that id is
	// missing, malformed or references a non existing resource, an error is
	// returned. If any of the Accounts are associated with another User an
	// error is returned.
	AddUserAccounts(context.Context, *AddUserAccountsRequest) (*AddUserAccountsResponse, error)
	// SetUserAccounts sets the Accounts associated with this User.
	// Any existing Accounts are removed if they are not included in this request.
	// The provided request must include the User ID which the Accounts will be
	// associated with.  Any Accounts not included in this request but previously
	// associated with this user will be disassociated.
	// If the User ID is missing, malformed or references a non existing resource,
	// an error is returned.
	// If any of the Accounts are associated with another User an error is returned.
	SetUserAccounts(context.Context, *SetUserAccountsRequest) (*SetUserAccountsResponse, error)
	// RemoveUserAccounts removes Accounts from the specified User.
	// The provided request must include the User id which the Accounts
	// will be removed from. If the provided Account ids is not associated with the
	// provided User, an error is returned.
	RemoveUserAccounts(context.Context, *RemoveUserAccountsRequest) (*RemoveUserAccountsResponse, error)
	// ListResolvableAliases returns a list of Aliases which point to a resource
	// for which the provided user id has some permission.
	// If missing or malformed an error is returned.
	ListResolvableAliases(context.Context, *ListResolvableAliasesRequest) (*ListResolvableAliasesResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) AddUserAccounts(context.Context, *AddUserAccountsRequest) (*AddUserAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserAccounts not implemented")
}
func (UnimplementedUserServiceServer) SetUserAccounts(context.Context, *SetUserAccountsRequest) (*SetUserAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserAccounts not implemented")
}
func (UnimplementedUserServiceServer) RemoveUserAccounts(context.Context, *RemoveUserAccountsRequest) (*RemoveUserAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserAccounts not implemented")
}
func (UnimplementedUserServiceServer) ListResolvableAliases(context.Context, *ListResolvableAliasesRequest) (*ListResolvableAliasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResolvableAliases not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddUserAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUserAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddUserAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUserAccounts(ctx, req.(*AddUserAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetUserAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetUserAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SetUserAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetUserAccounts(ctx, req.(*SetUserAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveUserAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveUserAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RemoveUserAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveUserAccounts(ctx, req.(*RemoveUserAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListResolvableAliases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResolvableAliasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListResolvableAliases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListResolvableAliases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListResolvableAliases(ctx, req.(*ListResolvableAliasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controller.api.services.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "AddUserAccounts",
			Handler:    _UserService_AddUserAccounts_Handler,
		},
		{
			MethodName: "SetUserAccounts",
			Handler:    _UserService_SetUserAccounts_Handler,
		},
		{
			MethodName: "RemoveUserAccounts",
			Handler:    _UserService_RemoveUserAccounts_Handler,
		},
		{
			MethodName: "ListResolvableAliases",
			Handler:    _UserService_ListResolvableAliases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/api/services/v1/user_service.proto",
}

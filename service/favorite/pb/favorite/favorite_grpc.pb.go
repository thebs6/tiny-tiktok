// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: favorite.proto

package favorite

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_FavoriteAction_FullMethodName = "/favorite.UserService/FavoriteAction"
	UserService_FavoriteList_FullMethodName   = "/favorite.UserService/FavoriteList"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error)
	FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error) {
	out := new(FavoriteActionResp)
	err := c.cc.Invoke(ctx, UserService_FavoriteAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	out := new(FavoriteListResp)
	err := c.cc.Invoke(ctx, UserService_FavoriteList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionResp, error)
	FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedUserServiceServer) FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FavoriteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FavoriteAction(ctx, req.(*FavoriteActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FavoriteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FavoriteList(ctx, req.(*FavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "favorite.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _UserService_FavoriteAction_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _UserService_FavoriteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
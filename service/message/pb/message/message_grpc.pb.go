// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc3
// source: message.proto

package message

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
	MessageService_Chat_FullMethodName   = "/douyin.extra.second.MessageService/chat"
	MessageService_Action_FullMethodName = "/douyin.extra.second.MessageService/action"
)

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	Chat(ctx context.Context, in *DouyinMessageChatRequest, opts ...grpc.CallOption) (*DouyinMessageChatResponse, error)
	Action(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Chat(ctx context.Context, in *DouyinMessageChatRequest, opts ...grpc.CallOption) (*DouyinMessageChatResponse, error) {
	out := new(DouyinMessageChatResponse)
	err := c.cc.Invoke(ctx, MessageService_Chat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) Action(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error) {
	out := new(DouyinRelationActionResponse)
	err := c.cc.Invoke(ctx, MessageService_Action_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	Chat(context.Context, *DouyinMessageChatRequest) (*DouyinMessageChatResponse, error)
	Action(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) Chat(context.Context, *DouyinMessageChatRequest) (*DouyinMessageChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedMessageServiceServer) Action(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinMessageChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_Chat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Chat(ctx, req.(*DouyinMessageChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_Action_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Action(ctx, req.(*DouyinRelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.extra.second.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "chat",
			Handler:    _MessageService_Chat_Handler,
		},
		{
			MethodName: "action",
			Handler:    _MessageService_Action_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

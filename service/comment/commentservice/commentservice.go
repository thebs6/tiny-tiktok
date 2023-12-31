// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package commentservice

import (
	"context"

	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment           = comment.Comment
	CommentActionReq  = comment.CommentActionReq
	CommentActionResp = comment.CommentActionResp
	CommentListReq    = comment.CommentListReq
	CommentListResp   = comment.CommentListResp
	User              = comment.User

	CommentService interface {
		CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error)
		CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
	}

	defaultCommentService struct {
		cli zrpc.Client
	}
)

func NewCommentService(cli zrpc.Client) CommentService {
	return &defaultCommentService{
		cli: cli,
	}
}

func (m *defaultCommentService) CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error) {
	client := comment.NewCommentServiceClient(m.cli.Conn())
	return client.CommentAction(ctx, in, opts...)
}

func (m *defaultCommentService) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	client := comment.NewCommentServiceClient(m.cli.Conn())
	return client.CommentList(ctx, in, opts...)
}

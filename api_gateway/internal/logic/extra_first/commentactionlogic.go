package extra_first

import (
	"context"
	"net/http"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionReq) (resp *types.CommentActionResp, err error) {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "user.rpc",
		},
	})
	client := comment.NewCommentServiceClient(conn.Conn())

	userid := l.ctx.Value("payload").(int64)

	respRpc, err := client.CommentAction(l.ctx, &comment.CommentActionReq{
		UserId:      userid,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
	})
	return &types.CommentActionResp{
		StatusCode: http.StatusOK,
		StatusMsg:  respRpc.StatusMsg,
		Comment: types.Comment{
			ID: resp.Comment.ID,
			User: types.User{
				ID:   resp.Comment.User.ID,
				Name: resp.Comment.User.Name,
			},
			Content: resp.Comment.Content,
		},
	}, err
}

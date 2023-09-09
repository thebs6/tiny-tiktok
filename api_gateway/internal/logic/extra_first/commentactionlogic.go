package extra_first

import (
	"context"
	"encoding/json"
	"net/http"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"

	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
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
	// conn := zrpc.MustNewClient(zrpc.RpcClientConf{
	// 	Etcd: discov.EtcdConf{
	// 		Hosts: []string{"etcd:2379"},
	// 		Key:   "comment.rpc",
	// 	},
	// })
	// client := comment.NewCommentServiceClient(conn.Conn())

	uid, err := l.ctx.Value("payload").(json.Number).Int64()
	if err != nil {
		logc.Info(l.ctx, "payload.(string) failed")
		return &types.CommentActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "fail!",
		}, nil
	}

	respRpc, err := l.svcCtx.CommentRpc.CommentAction(l.ctx, &comment.CommentActionReq{
		UserId:      uid,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	})
	if err != nil {
		logc.Alert(l.ctx, err.Error())
		resp = &types.CommentActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "fail!",
		}
		err = nil
		return
	}
	if req.ActionType == 1 {
		// publish comment
		resp = &types.CommentActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  respRpc.StatusMsg,
			Comment: types.Comment{
				ID: respRpc.Comment.Id,
				User: types.User{
					Id:   respRpc.Comment.User.Id,
					Name: respRpc.Comment.User.Name,
				},
				Content:    respRpc.Comment.Content,
				CreateDate: respRpc.Comment.CreateDate,
			},
		}
	} else {
		// delete common
		resp = &types.CommentActionResp{
			StatusCode: http.StatusOK,
			StatusMsg:  respRpc.StatusMsg,
			Comment:    types.Comment{},
		}
	}

	return
}

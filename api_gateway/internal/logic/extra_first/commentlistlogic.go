package extra_first

import (
	"context"
	"net/http"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"etcd:2379"},
			Key:   "comment.rpc",
		},
	})
	client := comment.NewCommentServiceClient(conn.Conn())

	// userid := l.ctx.Value("payload").(int64)
	respRpc, err := client.CommentList(l.ctx, &comment.CommentListReq{
		VideoId: req.VideoID,
	})
	if err != nil {
		resp = &types.CommentListResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Fail to get the comment list",
		}
		err = nil
		return
	}
	var comments []types.Comment
	for _, respComment := range respRpc.CommentList {
		comments = append(comments, types.Comment{
			Content:    respComment.Content,
			CreateDate: respComment.CreateDate,
			ID:         respComment.Id,
			User: types.User{
				Id:   respComment.User.Id,
				Name: respComment.User.Name,
			},
		})
	}

	resp = &types.CommentListResp{
		StatusCode:  http.StatusOK,
		StatusMsg:   respRpc.StatusMsg,
		CommentList: comments,
	}

	return
}

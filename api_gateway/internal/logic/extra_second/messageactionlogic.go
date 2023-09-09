package extra_second

import (
	"context"
	"encoding/json"
	"strconv"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/message/pb/message"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type MessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageActionLogic) MessageAction(req *types.MessageActionReq) (resp *types.MessageActionResp, err error) {
	// todo: add your logic here and delete this line
	//解析请求字段
	resp = &types.MessageActionResp{
		StatusCode: 0,
		StatusMsg:  "发送成功",
	}
	toUserId, err := strconv.ParseInt(req.ToUserID, 10, 64)
	if err != nil {
		logx.Errorf("to user id 解析错误 ： %s", err.Error())
		resp.StatusCode = 1
		resp.StatusMsg = "发送失败"
		return resp, err
	}

	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"etcd:2379"},
			Key:   "message.rpc",
		},
	})
	s := l.ctx.Value("payload").(json.Number).String()
	fromUserId, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		logx.Errorf("from user id 解析错误 ： %s", err.Error())
		resp.StatusCode = 1
		resp.StatusMsg = "发送失败"
		return resp, err
	}
	client := message.NewMessageServiceClient(conn.Conn())
	respRpc, err := client.Action(l.ctx, &message.DouyinRelationActionRequest{
		FromUserId: fromUserId,
		Token:      req.Token,
		ToUserId:   toUserId,
		ActionType: 1,
		Content:    req.Content,
	})
	if respRpc.StatusCode != 0 {
		resp.StatusCode = 1
		resp.StatusMsg = "发送失败"
		return resp, err
	}
	return resp, nil
}

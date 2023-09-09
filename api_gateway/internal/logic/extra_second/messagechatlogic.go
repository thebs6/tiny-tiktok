package extra_second

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/message/pb/message"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageChatLogic) MessageChat(req *types.MessageChatReq) (resp *types.MessageChatResp, err error) {
	// todo: add your logic here and delete this line
	id, err := strconv.ParseInt(req.ToUserID, 10, 64)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var Time int64
	if req.PreMsgTime == "" || req.PreMsgTime == "0" {
		Time = time.Now().Unix()
	}
	Time, _ = strconv.ParseInt(req.PreMsgTime, 10, 64)

	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"etcd:2379"},
			Key:   "message.rpc",
		},
	})

	client := message.NewMessageServiceClient(conn.Conn())
	respRpc, err := client.Chat(context.Background(), &message.DouyinMessageChatRequest{
		Token:      req.Token,
		ToUserId:   id,
		PreMsgTime: Time,
	})
	resp = &types.MessageChatResp{}
	err = copier.Copy(&resp, &respRpc)
	resp.StatusCode = "0"
	if err != nil {
		logx.Error(fmt.Sprintf("结构体复制错误+%s"), err)
		return nil, err
	}
	return resp, nil
}

package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"time"
	"tiny-tiktok/service/message/internal/svc"
	"tiny-tiktok/service/message/pb/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatLogic) Chat(in *message.DouyinMessageChatRequest) (*message.DouyinMessageChatResponse, error) {
	// todo: add your logic here and delete this line
	logx.Infof("%T", in.ToUserId)
	//将时间戳转换为时间
	t := time.Unix(in.PreMsgTime, 0)
	// 将时间格式化为字符串
	dateTime := t.Format("2006-01-02 15:04:05")

	resp, err := l.svcCtx.Model.FindList(l.ctx, in.ToUserId, dateTime)

	if err != nil {
		logx.Errorf("查找数据库失败 %s!", err.Error())
		return nil, err
	}
	messageList := make([]*message.Message, 10)
	err = copier.Copy(&messageList, resp)
	if err != nil {
		logx.Errorf("结构体复制错误%s!", err.Error())
		return nil, err
	}
	logx.Info(messageList[0].CreateTime)
	return &message.DouyinMessageChatResponse{
		StatusCode:  0,
		StatusMsg:   "返回成功",
		MessageList: messageList,
	}, nil
}

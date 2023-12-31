package lark

import (
	"context"
	"github.com/MR5356/notify"
	"github.com/go-lark/lark"
	"github.com/sirupsen/logrus"
)

type WebhookBot struct {
	bot *lark.Bot
}

func NewWebhookBot(hookURL string) *WebhookBot {
	return &WebhookBot{
		bot: lark.NewNotificationBot(hookURL),
	}
}

func (b *WebhookBot) Send(ctx context.Context, msg *notify.Message) error {
	var msgBuffer *lark.MsgBuffer
	switch msg.Type {
	case notify.MessageTypeCard:
		msgBuffer = cardBuilder(msg)
	default:
		msgBuffer = lark.NewMsgBuffer(lark.MsgText).Text(msg.Title)
	}
	v2, err := b.bot.PostNotificationV2(msgBuffer.Build())
	logrus.Debugf("post notification v2: %+v", v2)
	return err
}

func (b *WebhookBot) Name() string {
	return "lark"
}

func (b *WebhookBot) Params() []notify.Param {
	return []notify.Param{
		{
			Name: "hook_url",
			Desc: "飞书机器人webhook地址",
			Type: "string",
		},
	}
}

package mock

import (
	"context"
	"github.com/MR5356/notify"
	"github.com/MR5356/notify/utils"
	"github.com/sirupsen/logrus"
)

type Notifier struct {
}

func New() *Notifier {
	return &Notifier{}
}

func (n *Notifier) Send(ctx context.Context, msg *notify.Message) error {
	logrus.Info(utils.Struct2String(msg))
	return nil
}

func (n *Notifier) Name() string {
	return "mock"
}

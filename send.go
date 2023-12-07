package notify

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

var ErrSendNotification = errors.New("failed to send message")

func (n *Notify) Send(ctx context.Context, msg *Message) error {
	var eg errgroup.Group
	for _, notifier := range n.notifiers {
		if notifier == nil {
			continue
		}
		eg.Go(func() error {
			return notifier.Send(ctx, msg)
		})
	}
	err := eg.Wait()
	if err != nil {
		err = errors.Wrap(ErrSendNotification, err.Error())
	}
	return nil
}

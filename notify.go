package notify

import "context"

var ntf = New()

type Notify struct {
	notifiers []Notifier
}

type Notifier interface {
	Send(ctx context.Context, msg *Message) error
}

func New(ntfs ...Notifier) *Notify {
	return &Notify{
		notifiers: ntfs,
	}
}

func (n *Notify) WithNotifier(ntfs ...Notifier) *Notify {
	n.notifiers = append(n.notifiers, ntfs...)
	return n
}

func Default() *Notify {
	return ntf
}

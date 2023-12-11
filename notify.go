package notify

import "context"

var ntf = New()

type Notify struct {
	notifiers []Notifier
}

type Notifier interface {
	// Name 返回通知器的名称。
	//
	// 返回一个字符串类型。
	Name() string
	Send(ctx context.Context, msg *Message) error
	// Params 返回 New 函数需要的参数列表。
	//
	// 它不接受任何参数。
	// 返回类型是 []Param。
	Params() []Param
}

type Param struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Type string `json:"type"`
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

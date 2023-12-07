# Go Notify

 - [x] 支持发送飞书普通消息
 - [x] 支持发送飞书卡片消息
 - [ ] 支持发送email

## 使用说明
### 发送飞书普通消息
```go
package main

import (
	"context"
	"github.com/MR5356/notify"
	"github.com/MR5356/notify/notifier/lark"
)

func main() {
	msg := notify.NewMessage("test message")
	ntf := notify.New().WithNotifier(lark.NewWebhookBot("webhook-url"))
	err := ntf.Send(context.Background(), msg)
	if err != nil {
		panic(err)
	}
}
```

### 发送飞书卡片消息
```go
package main

import (
	"context"
	"github.com/MR5356/notify"
	"github.com/MR5356/notify/notifier/lark"
)

func main() {
	msg := notify.NewMessage("test message").
		CardBuilder().
		WithLevel(notify.MessageLevelError).
		AddText(notify.MessageLayoutDefault, true, "**时间**\n2022-01-01 00:00:00", "**级别**\ninfo").
		AddText(notify.MessageLayoutBisected, true, "**时间**\n2022-01-01 00:00:00", "**级别**\ninfo").
		AddButton(notify.MessageLayoutDefault, []*notify.Button{
			{
				Key:      "test",
				URL:      "https://docker.ac.cn",
				Level:    notify.MessageLevelInfo,
				Markdown: true,
			},
		}).
		AddImage(notify.MessageLayoutDefault, []*notify.Image{
			{
				Key:      "**时间**\n2022-01-01 00:00:00",
				Image:    "img_v2_bd87def3-f505-424d-a571-3f592a3b1b5g",
				Markdown: true,
			},
		}).
		AddImage(notify.MessageLayoutBisected, []*notify.Image{
			{
				Key:      "**时间**\n2022-01-01 00:00:00",
				Image:    "img_v2_bd87def3-f505-424d-a571-3f592a3b1b5g",
				Markdown: true,
			},
		}).WithNote("test note", true).
		Build()
	ntf := notify.New().WithNotifier(lark.NewWebhookBot("webhook-url"))
	err := ntf.Send(context.Background(), msg)
	if err != nil {
		panic(err)
	}
}
```
package lark

import (
	"github.com/MR5356/notify"
	"github.com/go-lark/lark"
	"github.com/go-lark/lark/card"
)

func cardBuilder(msg *notify.Message) *lark.MsgBuffer {
	builder := lark.NewCardBuilder()

	var body []card.Element

	// 处理正文
	for _, block := range msg.Body {
		var elements []*card.FieldBlock
		var bs []card.Element
		var div []card.Element
		for _, element := range block.Elements {
			if element.Type == notify.MessageElementTypeText {
				t := builder.Text(element.Key)
				if element.Markdown {
					t.LarkMd()
				}
				f := builder.Field(t)
				if block.Layout == notify.MessageLayoutBisected {
					f.Short()
				}
				elements = append(elements, f)
			}
			if element.Type == notify.MessageElementTypeButton {
				t := builder.Text(element.Key)
				if element.Markdown {
					t.LarkMd()
				}
				b := builder.Button(t).URL(element.Value)
				switch element.Level {
				case notify.MessageLevelError:
					b.Danger()
				case notify.MessageLevelWarning:
					b.Primary()
				default:
					b.Default()
				}
				bs = append(bs, b)
			}
			if element.Type == notify.MessageElementTypeImage {
				if block.Layout == notify.MessageLayoutBisected {
					t := builder.Text(element.Key)
					if element.Markdown {
						t.LarkMd()
					}
					div = append(div, builder.Div().Text(t).Extra(builder.Img(element.Value)))
				} else {
					div = append(div, builder.Img(element.Value))
				}

			}
		}
		if len(elements) > 0 {
			body = append(body, builder.Div(elements...))
		}
		if len(bs) > 0 {
			body = append(body, builder.Action(bs...))
		}
		if len(div) > 0 {
			body = append(body, div...)
		}
	}

	// 加入note
	if msg.Note != nil {
		if msg.Note.Markdown {
			body = append(body, builder.Note().AddText(builder.Text(msg.Note.Key).LarkMd()))
		} else {
			body = append(body, builder.Note().AddText(builder.Text(msg.Note.Key)))
		}
	}

	c := builder.Card(body...).Title(msg.Title)

	switch msg.Level {
	case notify.MessageLevelError:
		c.Carmine()
	case notify.MessageLevelWarning:
		c.Yellow()
	default:
		c.Blue()
	}

	return lark.NewMsgBuffer(lark.MsgInteractive).Card(c.String())
}

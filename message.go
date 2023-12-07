package notify

import (
	"github.com/mcuadros/go-defaults"
)

const (
	MessageLevelInfo    MessageLevel = "info"
	MessageLevelWarning MessageLevel = "warning"
	MessageLevelError   MessageLevel = "error"

	MessageLayoutDefault  MessageLayout = "default"
	MessageLayoutBisected MessageLayout = "bisected"

	MessageTypeText MessageType = "text"
	MessageTypeCard MessageType = "card"

	MessageElementTypeText   MessageElementType = "text"
	MessageElementTypeImage  MessageElementType = "image"
	MessageElementTypeButton MessageElementType = "button"
)

type MessageLevel string
type MessageLayout string
type MessageType string
type MessageElementType string

type Message struct {
	Title string          `json:"title"`
	Type  MessageType     `json:"type" default:"text"`
	Level MessageLevel    `json:"level" default:"info"`
	Note  *MessageElement `json:"note"`
	Body  MessageBody     `json:"body" default:"[]"`

	Receivers []*Receiver `json:"receivers" default:"[]"`
}

type MessageBody []*MessageBlock

type Receiver struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type MessageBlock struct {
	Layout   MessageLayout     `json:"layout" default:"default"`
	Elements []*MessageElement `json:"elements" default:"[]"`
}

type MessageElement struct {
	Type     MessageElementType `json:"type" default:"text"`
	Level    MessageLevel       `json:"level" default:"info"`
	Key      string             `json:"key"`
	Value    string             `json:"value"`
	Markdown bool               `json:"markdown"`
}

type MessageCard Message

func NewMessage(title string) *Message {
	message := new(Message)
	defaults.SetDefaults(message)
	message.Title = title
	return message
}

func (m *Message) AddReceiver(id, _type string) *Message {
	m.Receivers = append(m.Receivers, &Receiver{
		Id:   id,
		Type: _type,
	})
	return m
}

func (m *MessageCard) WithLevel(level MessageLevel) *MessageCard {
	m.Level = level
	return m
}

func (m *MessageCard) WithNote(note string, markdown bool) *MessageCard {
	m.Note = newMessageElement(MessageElementTypeText, MessageLevelInfo, note, "", markdown)
	return m
}

func (m *MessageCard) AddText(layout MessageLayout, markdown bool, texts ...string) *MessageCard {
	block := newMessageBlock(layout)
	for _, text := range texts {
		block.Elements = append(block.Elements, newMessageElement(MessageElementTypeText, MessageLevelInfo, text, "", markdown))
	}
	m.Body = append(m.Body, block)
	return m
}

type Button struct {
	Key      string       `json:"text"`
	URL      string       `json:"url"`
	Level    MessageLevel `json:"level"`
	Markdown bool         `json:"markdown"`
}

func (m *MessageCard) AddButton(layout MessageLayout, buttons []*Button) *MessageCard {
	block := newMessageBlock(layout)
	for _, v := range buttons {
		block.Elements = append(block.Elements, newMessageElement(MessageElementTypeButton, v.Level, v.Key, v.URL, v.Markdown))
	}
	m.Body = append(m.Body, block)
	return m
}

type Image struct {
	Key      string `json:"key"`
	Image    string `json:"image"`
	Markdown bool   `json:"markdown"`
}

func (m *MessageCard) AddImage(layout MessageLayout, images []*Image) *MessageCard {
	block := newMessageBlock(layout)
	for _, v := range images {
		block.Elements = append(block.Elements, newMessageElement(MessageElementTypeImage, MessageLevelInfo, v.Key, v.Image, v.Markdown))
	}
	m.Body = append(m.Body, block)
	return m
}

func (m *Message) CardBuilder() *MessageCard {
	m.Type = MessageTypeCard
	return (*MessageCard)(m)
}

func (m *MessageCard) Build() *Message {
	return (*Message)(m)
}

func newMessageBlock(layout MessageLayout) *MessageBlock {
	block := new(MessageBlock)
	defaults.SetDefaults(block)
	block.Layout = layout
	block.Elements = make([]*MessageElement, 0)
	return block
}

func newMessageElement(_type MessageElementType, level MessageLevel, key, value string, markdown bool) *MessageElement {
	element := new(MessageElement)
	defaults.SetDefaults(element)
	element.Type = _type
	element.Level = level
	element.Key = key
	element.Value = value
	element.Markdown = markdown
	return element
}

//func (m *MessageBody) ToHtml() string {
//	var res string
//	for _, block := range *m {
//		var style string
//		switch block.Layout {
//		case MessageLayoutBisected:
//			style = "display: flex; justify-content: space-between; gap: 10px; align-items: center"
//		}
//
//		div := fmt.Sprintf("<div style='%s'>", style)
//
//		for _, element := range block.Elements {
//			switch element.Type {
//			case MessageElementTypeText:
//				if element.Markdown {
//					div += fmt.Sprintf("<div><p>%s</p></div>", utils.Md2Html(element.Key))
//				} else {
//					div += fmt.Sprintf("<div><p>%s</p></div>", element.Key)
//				}
//			case MessageElementTypeButton:
//				if element.Markdown {
//					div += fmt.Sprintf("<div><a href='%s'>%s</a></div>", element.Value, utils.Md2Html(element.Key))
//				} else {
//					div += fmt.Sprintf("<div><a href='%s'>%s</a></div>", element.Value, element.Key)
//				}
//			case MessageElementTypeImage:
//				if block.Layout == MessageLayoutBisected {
//					div += fmt.Sprintf("<div><p>%s</p></div>", utils.Md2Html(element.Key))
//					div += fmt.Sprintf("<div><img src='%s' style='width: 10rem' alt='%s' /></div>", element.Value, element.Key)
//				} else {
//					div += fmt.Sprintf("<div><img src='%s' alt='%s' /></div>", element.Value, element.Key)
//				}
//			}
//		}
//		div += "</div>"
//		res += div
//	}
//	return res
//}

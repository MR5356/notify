package notify

import (
	"reflect"
	"testing"
)

func TestNewMessage(t *testing.T) {
	type args struct {
		title string
		body  []*MessageBlock
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{
			name: "test",
			args: args{
				title: "test1",
				body:  []*MessageBlock{},
			},
			want: &Message{
				Title: "test1",
				Level: "info",
				Body:  []*MessageBlock{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.title, tt.args.body...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

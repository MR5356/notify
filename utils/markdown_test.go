package utils

import "testing"

func TestMd2Html(t *testing.T) {
	type args struct {
		md string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				md: "**test** *aa*",
			},
			want: "<p><strong>test</strong> <em>aa</em></p>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md2Html(tt.args.md); got != tt.want {
				t.Errorf("Md2Html() = %v, want %v", got, tt.want)
			}
		})
	}
}

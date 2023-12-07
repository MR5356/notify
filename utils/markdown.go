package utils

import (
	"bytes"
	"github.com/yuin/goldmark"
	"strings"
)

func Md2Html(md string) string {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\n")
}

package server

import (
	"strings"
	"testing"
	"time"
)

func TestSlack_Send(t *testing.T) {
	s := NewSlack("your slack webhook url", "")
	s.Send("今天是个好日子")
}

func TestSlack_SendNews(t *testing.T) {
	g := NewGocnNew(nil)
	err, content := g.GetNewsContent(time.Now())
	if err != nil {
		t.Fatalf("err:%v\n", err)
	}
	s := NewSlack("your slack webhook url", "")
	c := strings.Join(content, "")
	s.Send(c)
}

package server

import (
	"strings"
	"testing"
	"time"
)

func TestWeCom_Send(t *testing.T) {
	w := NewWeCom("your bot key", "")
	w.Send("今天是个好日子, 消息忽略")
}

func TestWeCom_SendNews(t *testing.T) {
	w := NewWeCom("your bot key", "")
	g := NewGocnNew(nil)
	_, res := g.GetNewsContent(time.Now())
	r := strings.Join(res, "")
	w.Send(r)
}

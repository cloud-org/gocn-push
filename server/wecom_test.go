package server

import "testing"

func TestWeCom_Send(t *testing.T) {
	w := NewWeCom("your bot key", "")
	w.Send("今天是个好日子, 消息忽略")
}

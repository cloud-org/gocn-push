package server

import (
	"strings"
	"testing"
	"time"
)

func TestDingTalk_Send(t *testing.T) {
	d := NewDingTalk("you dingtalk bot token","")
	d.Send("今天是个好日子 yuque 语雀")
}

func TestDingTalk_SendNews(t *testing.T) {
	err, content := GetNewsContent(time.Now().Add(-24 * time.Hour))
	if err != nil {
		t.Fatalf("err:%v\n", err)
	}
	d := NewDingTalk("your dingtalk bot token", "")
	c := strings.Join(content, "")
	c = c + "\n from 语雀机器人"
	d.Send(c)
}

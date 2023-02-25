package server

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestGetNotify(t *testing.T) {
	token := "xxxx"
	n, _ := GetNotify("slack", token)
	err, contents := NewGoCnNew2023(nil).GetNewsContent(time.Now().Add(-24 * time.Hour))
	if err != nil {
		log.Printf("获取新闻发生错误, err: %v\n", err)
		return
	}
	content := strings.Join(contents, "")
	n.Send(content)
}

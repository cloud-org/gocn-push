package server

import (
	"fmt"
	"github.com/cloud-org/msgpush"
)

const timeFormat = "2006-01-02"

type TokenEnable struct {
	Enable bool   `json:"enable"`
	Token  string `json:"token"`
}

func GetNotify(software, token string) (msgpush.NotifyPush, error) {
	switch software {
	case "dingtalk": // 钉钉推送
		return msgpush.NewDingTalk(token), nil
	case "wecom": // 企业微信推送
		return msgpush.NewWeCom(token), nil
	case "slack": // slack 推送
		return msgpush.NewSlack(token), nil
	case "pushdeer":
		return msgpush.NewPushDeer(token), nil
	case "feishu":
		return msgpush.NewFeiShu(token), nil
	default:
		return nil, fmt.Errorf("暂时不支持类型 %v", software)
	}
}

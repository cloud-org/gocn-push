package server

import "fmt"

const timeFormat = "2006-01-02"

type NotifyPush interface {
	Send(string) error
	String() string
}

type TokenEnable struct {
	Enable bool   `json:"enable"`
	Token  string `json:"token"`
}

func GetNotify(software, token string) (NotifyPush, error) {
	switch software {
	case "dingtalk": // 钉钉推送
		return NewDingTalk(token, ""), nil
	case "wecom": // 企业微信推送
		return NewWeCom(token, ""), nil
	case "slack": // slack 推送
		return NewSlack(token, ""), nil
	case "pushdeer":
		return NewPushDeer(token), nil
	default:
		return nil, fmt.Errorf("暂时不支持类型 %v", software)
	}
}

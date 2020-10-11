package server

import (
	"fmt"
	"github.com/imroc/req"
)

const baseUrl = "https://oapi.dingtalk.com/robot/send?access_token="

type SendTextContent struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

type DingTalk struct {
	Token  string
	ReqUrl string
	Pre    string // 日期不同则可以进行抓取了
}

func NewDingTalk(token string, pre string) *DingTalk {
	return &DingTalk{Token: token, ReqUrl: baseUrl + token, Pre: pre}
}

func (d *DingTalk) Send(content string) error {
	return d.SendText(content)
}

func (d *DingTalk) SendText(content string) error {
	resp, err := req.Post(d.ReqUrl, req.BodyJSON(&SendTextContent{
		Msgtype: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}))
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp.String())
	return err
}

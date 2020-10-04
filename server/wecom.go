package server

import (
	"fmt"
	"github.com/imroc/req"
)

const wecomBase = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="

type weComSendTextContent struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

type WeCom struct {
	Token  string
	ReqUrl string
	Pre    string
	Flag   bool
}

//NewWeCom
func NewWeCom(token string) *WeCom {
	return &WeCom{Token: token, ReqUrl: wecomBase + token}
}

func (w *WeCom) Send(content string) error {
	return w.SendText(content)
}

func (w *WeCom) SendText(content string) error {
	resp, err := req.Post(w.ReqUrl, req.BodyJSON(&weComSendTextContent{
		Msgtype: "text",
		Text: struct {
			Content             string   `json:"content"`
			MentionedList       []string `json:"mentioned_list"`
			MentionedMobileList []string `json:"mentioned_mobile_list"`
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

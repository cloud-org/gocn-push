package server

import (
	"fmt"

	"github.com/imroc/req"
)

type slackTextContent struct {
	Text string `json:"text"`
}

type Slack struct {
	ReqUrl string
	Pre    string
}

func NewSlack(token string, pre string) *Slack {
	return &Slack{ReqUrl: token}
}

func (s *Slack) Send(content string) error {
	return s.sendText(content)
}

func (s *Slack) sendText(content string) error {
	resp, err := req.Post(s.ReqUrl, req.BodyJSON(&slackTextContent{Text: content}))
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp.String())
	return err
}

func (s *Slack) String() string {
	return "slack"
}

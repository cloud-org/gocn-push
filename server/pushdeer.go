package server

import (
	"fmt"
	"net/url"
	"time"

	"github.com/imroc/req"
)

type PushDeer struct {
	Token string
}

func NewPushDeer(token string) *PushDeer {
	return &PushDeer{Token: token}
}

func (p *PushDeer) Send(content string) error {
	return p.SendMd(content)
}

func (p *PushDeer) SendMd(content string) error {
	// https://api2.pushdeer.com/message/push?pushkey=<key>&text=标题&desp=<markdown>&type=markdown
	title := fmt.Sprintf("GoCN 每日新闻 (%s)", time.Now().Format(timeFormat))
	// urlencode
	reqUrl := fmt.Sprintf("https://api2.pushdeer.com/message/push?pushkey=%s&text=%s&desp=%s&type=markdown",
		p.Token, url.QueryEscape(title), url.QueryEscape(content))
	//log.Println("url is", reqUrl)
	resp, err := req.Post(reqUrl)
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp.String())
	return err
}

func (p *PushDeer) String() string {
	return "pushdeer"
}

package server

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/cloud-org/msgpush"

	"github.com/ronething/gocn-push/config"
)

//通用方法重构

type NewsPush struct {
	Pre     string
	Notifys []msgpush.NotifyPush
}

func (n *NewsPush) InitNotifys() error {
	notifyValue := config.Config.GetStringMap("notify")
	log.Printf("notifyValue is %v\n", notifyValue)
	for software, value := range notifyValue {
		data, err := json.Marshal(value)
		if err != nil {
			log.Printf("marshal err: %+v\n", err)
			return err
		}
		var tokenEnable TokenEnable
		if err = json.Unmarshal(data, &tokenEnable); err != nil {
			log.Printf("unmarshal err: %+v\n", err)
			return err
		}

		log.Printf("software: %s, data: %v\n", software, tokenEnable)
		if tokenEnable.Enable {
			notifyPush, err := GetNotify(software, tokenEnable.Token)
			if err != nil {
				log.Printf("get notify err: %v\n", err)
				return err
			}
			n.Notifys = append(n.Notifys, notifyPush)
		}
	}
	log.Printf("len(n.Notifys): %d res: %v\n", len(n.Notifys), n.Notifys)
	return nil
}

func (n *NewsPush) Push() {
	now := time.Now().Format(timeFormat)
	if n.Pre != now {
		err, contents := NewGoCnNew2023(nil).GetNewsContent(time.Now())
		if err != nil {
			log.Printf("获取新闻发生错误, err: %v\n", err)
			return
		}
		content := strings.Join(contents, "")
		for i := 0; i < len(n.Notifys); i++ {
			if n.Notifys[i].String() == "wecom" { // send text
				w := n.Notifys[i].(*msgpush.WeCom)
				if err = w.SendText(content); err != nil {
					log.Printf("wecom 推送失败, err: %v\n", err)
				}
				continue
			}
			if err = n.Notifys[i].Send(content); err != nil {
				log.Printf("%s 推送发生错误, err: %v\n", n.Notifys[i].String(), err)
				continue
			}
		}
		n.Pre = now
	}
}

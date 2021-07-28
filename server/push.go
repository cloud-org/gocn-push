package server

import (
	"log"
	"strings"
	"time"
)

//通用方法重构

func (d *DingTalk) NewsPushToDingTalk() {
	log.Printf("执行任务将 gocn 新闻推送到钉钉")
	now := time.Now().Format(timeFormat)
	log.Printf("dingtalk pre is %v, now is %v\n", d.Pre, now)
	if d.Pre != now { // 抓取
		err, contents := GetNewsContent(time.Now())
		if err != nil {
			log.Printf("获取新闻发生错误, err: %v\n", err)
			return
		}
		content := strings.Join(contents, "")
		if err = d.Send(content); err != nil {
			log.Printf("推送发生错误, err: %v\n", err)
			return
		}
		d.Pre = now
	}
	return
}

func (w *WeCom) NewsPushToWeCom() {
	log.Printf("执行任务将 gocn 新闻推送到企业微信")
	now := time.Now().Format(timeFormat)
	log.Printf("wecom pre is %v, now is %v\n", w.Pre, now)
	if w.Pre != now { // 抓取
		err, contents := GetNewsContent(time.Now())
		if err != nil {
			log.Printf("获取新闻发生错误, err: %v\n", err)
			return
		}
		content := strings.Join(contents, "")
		if err = w.Send(content); err != nil {
			log.Printf("推送发生错误, err: %v\n", err)
			return
		}
		w.Pre = now
	}
	return
}

func (s *Slack) NewsPushToSlack() {
	log.Printf("执行任务将 gocn 新闻推送到 Slack")
	now := time.Now().Format(timeFormat)
	log.Printf("slack pre is %v, now is %v\n", s.Pre, now)
	if s.Pre != now { // 抓取
		err, contents := GetNewsContent(time.Now())
		if err != nil {
			log.Printf("获取新闻发生错误, err: %v\n", err)
			return
		}
		content := strings.Join(contents, "")
		if err = s.Send(content); err != nil {
			log.Printf("推送发生错误, err: %v\n", err)
			return
		}
		s.Pre = now
	}
	return
}

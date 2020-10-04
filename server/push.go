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
	log.Printf("dingtalk pre is %v, flag is %v\n", d.Pre, d.Flag)
	if !d.Flag || d.Pre != now { // 抓取
		err, contents := GetNewsContent(time.Now())
		if err != nil {
			log.Printf("获取新闻发生错误, err: %v\n", err)
			d.Flag = false
			return
		}
		content := strings.Join(contents, "")
		if err = d.Send(content); err != nil {
			log.Printf("推送发生错误, err: %v\n", err)
			return
		}
		d.Flag = true
		d.Pre = now
	}
	return
}

func (w *WeCom) NewsPushToWeCom() {
	log.Printf("执行任务将 gocn 新闻推送到企业微信")
	now := time.Now().Format(timeFormat)
	log.Printf("wecom pre is %v, flag is %v\n", w.Pre, w.Flag)
	if !w.Flag || w.Pre != now { // 抓取
		err, contents := GetNewsContent(time.Now())
		if err != nil {
			log.Printf("获取新闻发生错误, err: %v\n", err)
			w.Flag = false
			return
		}
		content := strings.Join(contents, "")
		if err = w.Send(content); err != nil {
			log.Printf("推送发生错误, err: %v\n", err)
			return
		}
		w.Flag = true
		w.Pre = now
	}
	return
}

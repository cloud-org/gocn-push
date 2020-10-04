package server

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/ronething/gocn-push/config"
	"log"
	"os"
)

type Scheduler struct {
	C *cron.Cron
}

//NewScheduler 创建调度器
func NewScheduler() *Scheduler {
	optLogs := cron.WithLogger(
		cron.VerbosePrintfLogger(
			log.New(os.Stdout, "[Cron]: ", log.LstdFlags)))

	c := cron.New(optLogs)
	return &Scheduler{C: c}

}

func (s *Scheduler) Run() {
	s.C.Start()
}

func (s *Scheduler) InitJob() {

	log.Printf("dingtalk enable is %v\n", config.Config.GetBool("dingtalk.enable"))
	if config.Config.GetBool("dingtalk.enable") {
		d := NewDingTalk(config.Config.GetString("dingtalk.token"))
		_, err := s.C.AddFunc(config.Config.GetString("dingtalk.spec"), d.NewsPushToDingTalk)
		if err != nil {
			log.Printf("添加任务失败, err: %v\n", err)
			return
		}
	}

	log.Printf("wecom enable is %v\n", config.Config.GetBool("wecom.enable"))
	if config.Config.GetBool("wecom.enable") {
		w := NewWeCom(config.Config.GetString("wecom.token"))
		_, err := s.C.AddFunc(config.Config.GetString("wecom.spec"), w.NewsPushToWeCom)
		if err != nil {
			log.Printf("添加任务失败, err: %v\n", err)
			return
		}
	}
}

func (s *Scheduler) Stop() context.Context {
	return s.C.Stop()
}

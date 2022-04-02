package server

import (
	"context"
	"github.com/cloud-org/msgpush"
	"log"
	"os"

	"github.com/ronething/gocn-push/config"

	"github.com/robfig/cron/v3"
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

	cronSpec := config.Config.GetString("cron")
	pre := config.Config.GetString("pre")
	n := NewsPush{Pre: pre, Notifys: make([]msgpush.NotifyPush, 0)}
	//n := NewsPush{Pre: pre} 效果相同
	err := n.InitNotifys()
	if err != nil {
		log.Fatalf("init notifys err: %v\n", err)
	}

	_, err = s.C.AddFunc(cronSpec, n.Push)
	if err != nil {
		log.Fatalf("add cron job err: %v", err)
	}

}

func (s *Scheduler) Stop() context.Context {
	return s.C.Stop()
}

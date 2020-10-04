package server

import (
	"context"
	"github.com/robfig/cron/v3"
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

func (s *Scheduler) Stop() context.Context {
	return s.C.Stop()
}

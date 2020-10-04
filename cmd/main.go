package main

import (
	"flag"
	"fmt"
	"github.com/ronething/gocn-push/config"
	"github.com/ronething/gocn-push/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	filePath string // 配置文件路径
	help     bool   // 帮助
)

func usage() {
	fmt.Fprintf(os.Stdout, `gocn-push - gocn news push service
Usage: gocn [-h help] [-c ./config.yaml]
Options:
`)
	flag.PrintDefaults()
}

func main() {

	flag.StringVar(&filePath, "c", "./config.yaml", "配置文件所在")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = usage
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}

	// 设置配置文件和静态变量
	config.SetConfig(filePath)

	scheduler := server.NewScheduler()
	d := server.NewDingTalk(config.Config.GetString("dingtalk.token"))
	_, err := scheduler.C.AddFunc("* * * * *", d.NewsPushToDingTalk)
	if err != nil {
		log.Printf("添加任务失败, err: %v\n", err)
		return
	}
	w := server.NewWeCom(config.Config.GetString("wecom.token"))
	_, err = scheduler.C.AddFunc("* * * * *", w.NewsPushToWeCom)
	if err != nil {
		log.Printf("添加任务失败, err: %v\n", err)
		return
	}

	scheduler.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// DONE: 优雅关停
	for {
		s := <-c
		log.Printf("[main] 捕获信号 %s", s.String())
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			// 停止调度器 并等待正在 running 的任务执行结束 TODO: 有没有必要设置一个 timeout 假设一直不停止怎么办
			ctx := scheduler.Stop()
			timer := time.NewTimer(1 * time.Second)
			for {
				select {
				case s = <-c: // 再次接收到中断信号 则直接退出
					if s == syscall.SIGINT {
						log.Printf("[main] 再次接收到退出信号 %s", s.String())
						goto End
					}
				case <-ctx.Done():
					log.Printf("[main] 调度器所有任务执行完成")
					goto End
				case <-timer.C:
					log.Printf("[main] 调度器有任务正在执行中")
					timer.Reset(1 * time.Second)
				}
			}
		End:
			timer.Stop()
			return // 很重要 不然程序无法退出
		case syscall.SIGHUP:
			log.Printf("[main] 终端断开信号，忽略")
		default:
			log.Printf("[main] other signal")
		}
	}
}

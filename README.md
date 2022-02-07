### gocn-push

- push gocn news to dingtalk/wecom/slack

### ChangeLog

- 小重构 定时任务只需要一个
- 增加 spec、enable 选项
- 初版

### usage

- deploy

```bash
$ make deploy
```

- supervisor check

```bash
$ supervisorctl -c /etc/supervisor/supervisored.conf
gocnpush    RUNNING   pid 23569, uptime 290 days, 15:46:06
```

### TODO

- [x] 通知不应该这么做，spec 定时其实只需要一条即可，然后只需要获取一次新闻，通知方式通过列表获取，enable 为 true 的进行通知，
所以配置文件需要进行变更

> 目前的做法就是每新增一次，需要新增一次 spec cron 以及 newNotify func，比较不灵活

### Acknowledgement

- [news_watch_notice](https://github.com/Han-Ya-Jun/news_watch_notice)
- [gocn](https://github.com/georgehao/gocn)

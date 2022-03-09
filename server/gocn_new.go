package server

import (
	"fmt"
	"log"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"

	"github.com/go-resty/resty/v2"
)

type GocnNew struct {
	Client *resty.Client
}

func NewGocnNew(client *resty.Client) *GocnNew {
	if client == nil {
		client = resty.New()
	}
	return &GocnNew{Client: client}
}

var headers = map[string]string{
	"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
}

func (g *GocnNew) GetNewsContent(publishTime time.Time) (error, []string) {
	// 获取列表 查看是否有指定日期的每日新闻 拿到对应的 topic
	resp, err := g.Client.R().SetQueryParams(map[string]string{
		"currentPage": "1",
		"cate2Id":     "18",
		"grade":       "new",
	}).SetResult(&TopicListResp{}).
		SetHeaders(headers).
		Get("https://gocn.vip/apiv3/topic/list")
	if err != nil {
		log.Printf("request topic list err: %v", err)
		log.Printf("err resp: %v", resp.String())
		return err, nil
	}

	v, ok := resp.Result().(*TopicListResp)
	if !ok {
		log.Printf("TopicListResp 断言失败")
		return fmt.Errorf("topic list resp assert err"), nil
	}

	if len(v.Data.List) == 0 {
		return fmt.Errorf("list len is 0"), nil
	}

	newTopic := v.Data.List[0]
	// 判断 title 是否包含指定日期
	if !g.containsDate(publishTime, newTopic.Title) {
		return fmt.Errorf("新闻未更新"), nil
	}

	// 如果获取到对应的 topic，则可以进行主题的获取
	topicId := newTopic.GUID
	// https://gocn.vip/api/topic/20991/info
	// apiv2
	resp, err = g.Client.R().
		SetResult(&TopicInfoResp{}).
		SetHeaders(headers).
		Get(fmt.Sprintf("https://gocn.vip/apiv3/topic/%s/info", topicId))
	if err != nil {
		log.Printf("request topic info err: %v", err)
		log.Printf("err resp: %v", resp.String())
		return err, nil
	}

	value, ok := resp.Result().(*TopicInfoResp)
	if !ok {
		log.Printf("TopicInfoResp 断言失败")
		return fmt.Errorf("topic info resp assert err"), nil
	}

	info := value.Data.Topic.ContentHTML

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(info)
	if err != nil {
		// 转换失败直接赋值
		log.Printf("convert err: %v\n", err)
		markdown = info
	}

	return nil, []string{markdown}
}

func (g *GocnNew) containsDate(publishTime time.Time, title string) bool {

	data := publishTime.Format("2006-01-02")
	dateOther := publishTime.Format("2006-01-2")
	dateOtherOther := publishTime.Format("2006-1-2")
	if strings.Contains(title, data) || strings.Contains(title, dateOther) || strings.Contains(title, dateOtherOther) {
		return true
	}

	return false
}

package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type GoCnNew2023 struct {
	Client *resty.Client
}

func NewGoCnNew2023(client *resty.Client) *GoCnNew2023 {
	if client == nil {
		client = resty.New()
	}
	return &GoCnNew2023{Client: client}
}

type GoCnNewsData2023 struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []NewTopic2023 `json:"list"`
	} `json:"data"`
}

type NewTopic2023 struct {
	GUID      string `json:"guid"`
	Name      string `json:"name"` // title
	UID       int    `json:"uid"`
	Ctime     int    `json:"ctime"`
	CntView   int    `json:"cntView"`
	CmtGUID   string `json:"cmtGuid"`
	SpaceGUID string `json:"spaceGuid"`
	Format    int    `json:"format"`
	Content   string `json:"content"` // 正文
}

var headers2023 = map[string]string{
	"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
	"refer":      "https://gocn.vip/c/3lQ6GbD5ny/s/Gd7BTB",
}

func (g *GoCnNew2023) GetNewsContent(publishTime time.Time) (error, []string) {
	// 获取列表 查看是否有指定日期的每日新闻 拿到对应的 topic
	resp, err := g.Client.R().SetQueryParams(map[string]string{
		"spaceGuid":   "Gd7BTB",
		"currentPage": "1",
		"sort":        "1",
	}).SetResult(&GoCnNewsData2023{}).
		SetHeaders(headers2023).
		Get("https://gocn.vip/api/files")
	if err != nil {
		log.Printf("request files list err: %v", err)
		log.Printf("err resp: %v", resp.String())
		return err, nil
	}

	v, ok := resp.Result().(*GoCnNewsData2023)
	if !ok {
		log.Printf("GoCnNewsData2023 断言失败")
		return fmt.Errorf("request files list resp assert err"), nil
	}

	if len(v.Data.List) == 0 {
		return fmt.Errorf("files list len is 0"), nil
	}

	// 遍历一整个列表进行日期判定
	newTopic, ok := g.getTopic(v.Data.List, publishTime)
	if !ok {
		return fmt.Errorf("新闻未更新"), nil
	}

	news, err := g.parseContent(newTopic.Name, newTopic.Content)
	if err != nil {
		log.Printf("解析 content 失败 %v\n", err)
		return err, nil
	}

	return nil, []string{*news}
}

func (g *GoCnNew2023) getTopic(topics []NewTopic2023, publishTime time.Time) (*NewTopic2023, bool) {
	for i := 0; i < len(topics); i++ {
		newTopic := topics[i]
		// 判断 title 是否包含指定日期
		if g.containsDate(publishTime, newTopic.Name) {
			return &newTopic, true
		}
	}

	return nil, false
}

func (g *GoCnNew2023) containsDate(publishTime time.Time, title string) bool {

	data := publishTime.Format("2006-01-02")
	dateOther := publishTime.Format("2006-01-2")
	dateOtherOther := publishTime.Format("2006-1-2")
	if strings.Contains(title, data) || strings.Contains(title, dateOther) || strings.Contains(title, dateOtherOther) {
		return true
	}

	return false
}

func (g *GoCnNew2023) parseContent(title string, data string) (*string, error) {
	var nodes []*Node
	err := json.Unmarshal([]byte(data), &nodes)
	if err != nil {
		log.Printf("unmarshal content err: %v", err)
		return nil, err
	}

	var news bytes.Buffer
	news.WriteString(title + "\n\n")
	for _, node := range nodes {
		nodeTexts := getText(node)
		if len(nodeTexts) <= 1 { // 长度为 0 或者是标题
			continue
		}
		if len(nodeTexts)&1 == 0 && len(nodeTexts) >= 8 { // 偶数 正文
			count := 1
			for i := 0; i < len(nodeTexts); i = i + 2 {
				news.WriteString(fmt.Sprintf("%d. %s %s\n", count, nodeTexts[i], nodeTexts[i+1]))
				count++
			}
			news.WriteString("\n")
		}
		if len(nodeTexts)&1 == 0 && len(nodeTexts) <= 4 { // 偶数 可能是宣传链接之类
			for i := 0; i < len(nodeTexts); i = i + 2 {
				news.WriteString(fmt.Sprintf("%s %s\n", nodeTexts[i], nodeTexts[i+1]))
			}
			news.WriteString("\n")
		}
		if len(nodeTexts)&1 == 1 { //奇数 编辑信息
			news.WriteString(nodeTexts[0] + "\n")
			// 适配奇怪的渲染 有些是 text 和 url 分开 有些则不是
			if len(nodeTexts) <= 3 {
				for i := 1; i < len(nodeTexts); i++ {
					news.WriteString(fmt.Sprintf("%s\n", nodeTexts[i]))
				}
			}
			if len(nodeTexts) > 3 {
				for i := 1; i < len(nodeTexts); i = i + 2 {
					news.WriteString(fmt.Sprintf("%s %s\n", nodeTexts[i], nodeTexts[i+1]))
				}
			}
		}
	}

	res := news.String()

	return &res, nil
}

type Node struct {
	Type     string  `json:"type"`
	Children []*Node `json:"children"`
	Text     string  `json:"text"`
}

//getText 递归获取 texts
func getText(node *Node) []string {
	var texts []string
	if node.Text != "" {
		texts = append(texts, node.Text)
	}

	for _, child := range node.Children {
		childTexts := getText(child)
		texts = append(texts, childTexts...)
	}

	return texts
}

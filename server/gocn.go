package server

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// gocn news
func GetNewsContent(publishTime time.Time) (e error, content []string) {
	var baseUrl string
	c := colly.NewCollector()
	//t:=time.Now().Add(-time.Hour*time.Duration(24))
	data := publishTime.Format("2006-01-02")
	dateOther := publishTime.Format("2006-01-2")
	// Find and visit all links
	c.OnHTML("div.title.media-heading > a", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, data) {
			baseUrl = "https://gocn.vip" + e.Attr("href")
			fmt.Printf("Link found: %q -> %s\n", e.Text, baseUrl)
		} else if strings.Contains(e.Text, dateOther) {
			baseUrl = e.Attr("href")
			fmt.Printf("Link found: %q -> %s\n", e.Text, baseUrl)
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	e = c.Visit("https://gocn.vip/topics/node18")

	if e != nil {
		return
	}
	if baseUrl == "" {
		return errors.New("news not update"), nil
	}
	b := colly.NewCollector()

	// Find and visit all links
	b.OnHTML("div.card-body.markdown.markdown-toc", func(e *colly.HTMLElement) {
		e.ForEach("ol > li", func(i int, e *colly.HTMLElement) {
			t := strings.Trim(e.Text, "\n")
			content = append(content, fmt.Sprintf("%v. ", i+1)+t+"\n")
		})
		e.ForEach("p,ul", func(i int, element *colly.HTMLElement) {
			authorIndex1 := strings.Index(element.Text, "编辑:")
			authorIndex2 := strings.Index(element.Text, "编辑：")
			// log.Printf("index1 is %v, index2 is %v\n", authorIndex1, authorIndex2)
			if authorIndex1 >= 0 || authorIndex2 >= 0 {
				// DONE: trim \n 然后自己再手动控制回车
				t := strings.Trim(element.Text, "\n")
				t = strings.Replace(t, "\n\n", "\n", -1) // 兼容行与行之间有两个回车的情况
				content = append(content, "\n"+t)
			} else { // no col the footer
				//content = append([]string{element.Text + "\n\n"}, content...)
			}
		})
	})

	b.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	e = b.Visit(baseUrl)
	if e != nil {
		return
	}
	// 补齐表头
	if len(content) > 0 {
		content = append([]string{fmt.Sprintf("GoCN 每日新闻 (%s)", data) + "\n\n"}, content...)
	}
	return
}

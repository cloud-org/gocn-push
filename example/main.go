package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Type     string  `json:"type"`
	Children []*Node `json:"children"`
	Text     string  `json:"text"`
}

func main() {
	data := `[{"type":"h2","lineHeight":"1.225","align":"start","children":[{"text":"GoCN 每日新闻 (2023-02-22)"}]},
	         {"type":"ol","children":[{"type":"li","children":[{"type":"lic","listStyleType":"","indent":0,"children":[{"text":"GO 语言依赖注入及相关框架 "},{"type":"a","url":"https://qiankunli.github.io/2023/02/16/go_ioc.html","target":"_blank","children":[{"text":"https://qiankunli.github.io/2023/02/16/go_ioc.html"}]},{"text":""}]}]},
	         {"type":"li","children":[{"type":"lic","listStyleType":"","indent":0,"children":[{"text":"一个简单可爱的Go 日志库 "},{"type":"a","url":"https://github.com/charmbracelet/log","target":"_blank","children":[{"text":"https://github.com/charmbracelet/log"}]},{"text":""}]}]}]},
	         {"type":"ul","children":[{"type":"li","children":[{"type":"lic","listStyleType":"","indent":0,"children":[{"text":"编辑: 鹿沐"}]}]},
	         {"type":"li","children":[{"type":"lic","listStyleType":"","indent":0,"children":[{"text":"订阅新闻: "},{"type":"a","url":"http://tinyletter.com/gocn","target":"_blank","children":[{"text":"http://tinyletter.com/gocn"}]},{"text":""}]}]},
	         {"type":"li","children":[{"type":"lic","listStyleType":"","indent":0,"children":[{"text":"招聘专区: "},{"type":"a","url":"https://gocn.vip/jobs","target":"_blank","children":[{"text":"https://gocn.vip/jobs"}]},{"text":""}]}]}]}]`
	//data := `[{"type":"ol","children":[{"type":"li","children":[{"children":[{"text":"了解 Go 中的指针的一页  "},{"type":"a","url":"https://medium.com/@Lekia/a-one-pager-to-understanding-pointers-in-go-ad6cbfac3afc","children":[{"text":"https://medium.com/@Lekia/a-one-pager-to-understanding-pointers-in-go-ad6cbfac3afc"}]},{"text":""}],"type":"lic"}]},{"type":"li","children":[{"type":"lic","children":[{"text":"GO-select 的实现原理 "},{"type":"a","url":"https://juejin.cn/post/7201423410168741946","children":[{"text":"https://juejin.cn/post/7201423410168741946"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"Golang：使用同步包将性能提高 10 倍并减少内存占用 "},{"type":"a","url":"https://medium.com/@aryehlevklein/golang-using-sync-package-to-10x-performance-and-reduce-memory-footprint-a1ed4ee14931","children":[{"text":"https://medium.com/@aryehlevklein/golang-using-sync-package-to-10x-performance-and-reduce-memory-footprint-a1ed4ee14931"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"云原生系列Go语言篇-错误处理 "},{"type":"a","url":"https://juejin.cn/post/7201509055713427513","children":[{"text":"https://juejin.cn/post/7201509055713427513"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"终极 2023 Web 服务器基准测试：NodeJS vs Java vs Rust vs Go "},{"type":"a","url":"https://medium.com/@alexeynovikov_89393/ultimate-2023-web-server-benchmark-nodejs-vs-java-vs-rust-vs-go-e367d932f699","children":[{"text":"https://medium.com/@alexeynovikov_89393/ultimate-2023-web-server-benchmark-nodejs-vs-java-vs-rust-vs-go-e367d932f699"}]},{"text":"","strikethrough":true}]}]}]},{"type":"p","children":[{"text":""}]},{"type":"ul","children":[{"type":"li","children":[{"type":"lic","children":[{"text":"编辑: zsr228"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"订阅新闻: http://tinyletter.com/gocn"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"招聘专区: https://gocn.vip/jobs"}]}]}]}]`

	var nodes []*Node
	err := json.Unmarshal([]byte(data), &nodes)
	if err != nil {
		panic(err)
	}

	var texts []string
	for _, node := range nodes {
		nodeTexts := getText(node)
		fmt.Println("node texts is ", nodeTexts, len(nodeTexts))
		texts = append(texts, nodeTexts...)
	}

	fmt.Println(texts)
}

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

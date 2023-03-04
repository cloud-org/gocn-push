package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
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
	//data = `[{"type":"ol","children":[{"type":"li","children":[{"children":[{"text":"了解 Go 中的指针的一页  "},{"type":"a","url":"https://medium.com/@Lekia/a-one-pager-to-understanding-pointers-in-go-ad6cbfac3afc","children":[{"text":"https://medium.com/@Lekia/a-one-pager-to-understanding-pointers-in-go-ad6cbfac3afc"}]},{"text":""}],"type":"lic"}]},{"type":"li","children":[{"type":"lic","children":[{"text":"GO-select 的实现原理 "},{"type":"a","url":"https://juejin.cn/post/7201423410168741946","children":[{"text":"https://juejin.cn/post/7201423410168741946"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"Golang：使用同步包将性能提高 10 倍并减少内存占用 "},{"type":"a","url":"https://medium.com/@aryehlevklein/golang-using-sync-package-to-10x-performance-and-reduce-memory-footprint-a1ed4ee14931","children":[{"text":"https://medium.com/@aryehlevklein/golang-using-sync-package-to-10x-performance-and-reduce-memory-footprint-a1ed4ee14931"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"云原生系列Go语言篇-错误处理 "},{"type":"a","url":"https://juejin.cn/post/7201509055713427513","children":[{"text":"https://juejin.cn/post/7201509055713427513"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"终极 2023 Web 服务器基准测试：NodeJS vs Java vs Rust vs Go "},{"type":"a","url":"https://medium.com/@alexeynovikov_89393/ultimate-2023-web-server-benchmark-nodejs-vs-java-vs-rust-vs-go-e367d932f699","children":[{"text":"https://medium.com/@alexeynovikov_89393/ultimate-2023-web-server-benchmark-nodejs-vs-java-vs-rust-vs-go-e367d932f699"}]},{"text":"","strikethrough":true}]}]}]},{"type":"p","children":[{"text":""}]},{"type":"ul","children":[{"type":"li","children":[{"type":"lic","children":[{"text":"编辑: zsr228"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"订阅新闻: http://tinyletter.com/gocn"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"招聘专区: https://gocn.vip/jobs"}]}]}]}]`
	//data = `[{"type":"ol","children":[{"type":"li","children":[{"children":[{"text":"ServiceWeaver：一个编写分布式应用程序的框架 https://opensource.googleblog.com/2023/03/introducing-service-weaver-framework-for-writing-distributed-applications.html"}],"type":"lic"}]},{"type":"li","children":[{"type":"lic","children":[{"text":"Go1.20 arena 能手动管理内存了，怎么用？ https://mp.weixin.qq.com/s/mwWMOwLsiY8EtODpyEoTIg"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"Go 语言性能剖析利器：pprof 实战 https://toutiao.io/posts/ye9g2eb"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"Go中gin框架中Session详解 https://juejin.cn/post/7205016004925423653"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"Go-Benchmark入门-基础篇（上） "},{"type":"a","url":"https://juejin.cn/post/7205764215222403132","children":[{"text":"https://juejin.cn/post/7205764215222403132"}]},{"text":""}]}]}]},{"type":"p","children":[{"text":""}]},{"type":"ul","children":[{"type":"li","children":[{"type":"lic","children":[{"text":"编辑: flint"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"订阅新闻: http://tinyletter.com/gocn"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"招聘专区: https://gocn.vip/jobs"}]}]}]}]`
	//data = `[{"type":"code_block","children":[{"type":"code_line","children":[{"text":"## GoCN 每日新闻 (2023-02-28)"}]},{"type":"code_line","children":[{"text":"1. [K8S] client-go 的正确打开方式 https://juejin.cn/post/7203690731276517432"}]},{"type":"code_line","children":[{"text":"2. 优化time.After后，性能提升34%，内存减少67% https://juejin.cn/post/7203274235426324536"}]},{"type":"code_line","children":[{"text":"3. 如何在Golang中进行热重载和优雅的关闭 https://medium.com/@ramseyjiang_22278/how-to-do-hot-reload-and-graceful-shutdown-in-golang-3f84a9f17d79"}]},{"type":"code_line","children":[{"text":"4. 带单位测试的Golang的代理设计模式 https://medium.com/@ramseyjiang_22278/proxy-design-pattern-in-golang-with-unit-tests-d9c6c0d01d49"}]},{"type":"code_line","children":[{"text":"5. 跨平台的Golang GUI库, 核心绑定自 Lazarus 创建的通用跨平台GUI库 liblcl https://github.com/ying32/govcl"}]},{"type":"code_line","children":[{"text":""}]},{"type":"code_line","children":[{"text":"- 编辑: Rolle"}]},{"type":"code_line","children":[{"text":"- 订阅新闻: http://tinyletter.com/gocn"}]},{"type":"code_line","children":[{"text":"- 招聘专区: https://gocn.vip/jobs"}]},{"type":"code_line","children":[{"text":""}]}],"lang":"markdown"},{"type":"p","children":[{"text":""}]}]`
	data = `[{"type":"h2","children":[{"text":"GoCN 每日新闻 (2023-02-26)"}]},{"type":"ol","children":[{"type":"li","children":[{"type":"lic","children":[{"text":"内存对齐这个事儿只能自己搞明白 "},{"type":"a","url":"https://mp.weixin.qq.com/s/TphkrCag_Wr6uD9thPMwcg","children":[{"text":"https://mp.weixin.qq.com/s/TphkrCag_Wr6uD9thPMwcg"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"紧急下班：修炼内功---内存模型和垃圾回收，不焦虑打工指南 "},{"type":"a","url":"https://mp.weixin.qq.com/s/qF5eiWYkcwTnZ2tFwpzx2Q","children":[{"text":"https://mp.weixin.qq.com/s/qF5eiWYkcwTnZ2tFwpzx2Q"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"Go 语言跨平台文件监听库 fsnotify 怎么使用？ "},{"type":"a","url":"https://mp.weixin.qq.com/s/tJ1LvDf14EKg-qQlJUQapQ","children":[{"text":"https://mp.weixin.qq.com/s/tJ1LvDf14EKg-qQlJUQapQ"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"解析 Golang 网络 IO 模型之 EPOLL "},{"type":"a","url":"https://mp.weixin.qq.com/s/xt0Elppc_OaDFnTI_tW3hg","children":[{"text":"https://mp.weixin.qq.com/s/xt0Elppc_OaDFnTI_tW3hg"}]},{"text":""}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"Golang HTTP 标准库实现原理 "},{"type":"a","url":"https://mp.weixin.qq.com/s/zFG6_o0IKjXh4RxKmPTt4g","children":[{"text":"https://mp.weixin.qq.com/s/zFG6_o0IKjXh4RxKmPTt4g"}]},{"text":""}]}]}]},{"type":"p","children":[{"text":""}]},{"type":"ul","children":[{"type":"li","children":[{"type":"lic","children":[{"text":"编辑: 鱼雷"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"订阅新闻: http://tinyletter.com/gocn"}]}]},{"type":"li","children":[{"type":"lic","children":[{"text":"招聘专区: https://gocn.vip/jobs"}]}]}]}]`

	var nodes []*Node
	err := json.Unmarshal([]byte(data), &nodes)
	if err != nil {
		panic(err)
	}

	var texts []string
	for _, node := range nodes {
		nodeTexts := getText(node)
		//fmt.Println("node texts is ", nodeTexts, len(nodeTexts))
		texts = append(texts, nodeTexts...)
	}

	//fmt.Println(texts)
	info := hello(texts)

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(info)
	if err != nil {
		// 转换失败直接赋值
		log.Printf("convert err: %v\n", err)
		markdown = info
	}

	fmt.Println(markdown)

}

func getText(node *Node) []string {
	var texts []string
	if node.Type != "" {
		texts = append(texts, node.Type)
	}
	text := strings.TrimSpace(strings.Trim(node.Text, "\n"))
	if text != "" {
		texts = append(texts, text)
	}

	for _, child := range node.Children {
		childTexts := getText(child)
		texts = append(texts, childTexts...)
	}

	return texts
}

func hello(data []string) string {

	var result strings.Builder
	var stack []string
	var stackText []string

	for i := 0; i < len(data); i++ {
		item := data[i]
		switch item {
		case "li", "ol", "ul", "h1", "h2":
			if len(stack) > 0 && len(stackText) > 0 {
				result.WriteString(fmt.Sprintf("</%s>\n", stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
				stackText = stackText[:len(stackText)-1]
			}
			result.WriteString(fmt.Sprintf("<%s>", item))
			stack = append(stack, item)
		case "lic", "p", "code_line", "code_block": // pass
		case "a":
			result.WriteString(fmt.Sprintf("<a href=\"%s\">%s</a>", data[i+1], data[i+1]))
			i = i + 1
		default:
			result.WriteString(item)
			stackText = append(stackText, item)
		}
	}

	//fmt.Println(stackText)
	for i := len(stack) - 1; i >= 0; i-- {
		result.WriteString(fmt.Sprintf("</%s>", stack[i]))
	}

	//fmt.Println(result.String())

	return result.String()
}

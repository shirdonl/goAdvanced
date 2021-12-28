//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

var tokens = make(chan struct{}, 20)

func Crawling(url string) []string {
	//打印URL
	fmt.Println(url)
	tokens <- struct{}{} //获取令牌
	list, err := Extracting(url)
	<-tokens //释放令牌
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	//创建通道数组
	worklist := make(chan []string)
	var n int
	n++ //n用来记录任务列表中的任务个数
	go func() { worklist <- os.Args[1:] }()
	//已经获取的链接
	picked := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !picked[link] {
				picked[link] = true
				//任务个数加1
				n++
				//开启goroutine
				go func(link string) {
					worklist <- Crawling(link)
				}(link)
			}
		}
	}
}

// 提取对指定的URL发出HTTP GET请求，进行解析并响应为HTML，并返回HTML文档中的链接
func Extracting(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

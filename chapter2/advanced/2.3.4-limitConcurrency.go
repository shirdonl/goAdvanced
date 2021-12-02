package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

//爬取URL
func Crawling(url string) []string {
	fmt.Println(url)
	list, err := Extracting(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	//创建通道数组
	worklist := make(chan []string)
	go func() { worklist <- os.Args[1:] }()

	//已经获取的链接
	picked := make(map[string]bool)
	for list := range worklist {
		for _, url := range list {
			if !picked[url] {
				picked[url] = true
				//开启goroutine
				go func(u string) {
					worklist <- Crawling(url)
				}(url)
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

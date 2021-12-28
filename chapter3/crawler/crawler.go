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

package crawler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Crawler struct {
}

func Get(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//读取网页的body内容
	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}
		result += string(buf[:n])
	}
	return
}

//将所有的网页内容爬取下来
func SpiderPage(url string) {
	fmt.Printf("正在爬取%s\n", url)
	//爬,将所有的网页内容爬取下来
	result, err := Get(url)
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}
	//把内容写入到文件
	filename := "page" + url + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("os.Create err = ", err1)
		return
	}
	//写内容
	f.WriteString(result)
	//关闭文件
	f.Close()
}

func Do(urls []string) {
	//因为很有可能爬虫还没有结束下面的循环就已经结束了，所以这里就需要且到通道
	for _, url := range urls {
		//var url string
		//将page阻塞
		go SpiderPage(url)
	}
}

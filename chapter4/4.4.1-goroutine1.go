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
	"io/ioutil"
	"net/http"
	"time"
)

//屏障模式响应结构体
type BarrierResponse struct {
	Err    error
	Resp   string
	Status int
}

// 构造请求
func doRequest(out chan<- BarrierResponse, url string) {
	res := BarrierResponse{}

	//设置HTTP客户端
	client := http.Client{
		Timeout: time.Duration(20 * time.Second),
	}

	//执行GET请求
	resp, err := client.Get(url)
	if resp != nil {
		res.Status = resp.StatusCode
	}
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	//将获取的结果数据放入通道
	out <- res
}

// 合并结果
func Barrier(urls ...string) {
	requestNumber := len(urls)

	in := make(chan BarrierResponse, requestNumber)
	response := make([]BarrierResponse, requestNumber)

	defer close(in)

	for _, urls := range urls {
		go doRequest(in, urls)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err, resp.Status)
			hasError = true
		}
		response[i] = resp
	}
	if !hasError {
		for _, resp := range response {
			fmt.Println(resp.Status)
		}
	}
}

func main() {
	//执行请求
	Barrier([]string{"https://www.baidu.com/",
		"https://www.weibo.com",
		"https://www.shirdon.com/"}...)
}

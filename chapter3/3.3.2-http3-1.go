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
	"flag"
	"fmt"
	"github.com/lucas-clemente/quic-go/http3"
	"log"
	"net/http"
)

func main() {
	//命令行参数
	var (
		port     = flag.Uint("p", 8080, "front port")
		useHttp3 = flag.Bool("enable_http3", false, "use http3")
		//certPath       = flag.String("c", "./testdata/cert.pem", "cert file path")
		//keyPath        = flag.String("k", "./testdata/priv.key", "key file path")
	)
	flag.Parse()

	//代理服务器设置
	proxyServer := http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", *port),
	}

	if *useHttp3 {
		fmt.Println("http3")
		http.HandleFunc("/", handleFunc)
		http3.ListenAndServeQUIC("localhost:6060", "./testdata/cert.pem", "./testdata/priv.key", nil)

		//if err := http3.ListenAndServe(proxyServer.Addr, *certPath, *keyPath, proxyServer.Handler); err != nil {
		//	log.Fatalln(err)
		//}
	} else {
		fmt.Println("http")
		if err := proxyServer.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi http3")
}

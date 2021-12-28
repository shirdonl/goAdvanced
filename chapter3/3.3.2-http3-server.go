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
	"net/http"
	"strings"
	"sync"

	_ "net/http/pprof"

	"gitee.com/shirdonl/goAdvanced/chapter3/testdata"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

type binds []string

func (b binds) String() string {
	return strings.Join(b, ",")
}

func (b *binds) Set(v string) error {
	*b = strings.Split(v, ",")
	return nil
}

func setupHandler(www string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is http3 root")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi,http3")
	})

	return mux
}

func main() {
	//绑定https://localhost:8088
	bs := binds{"localhost:8088"}
	//配置QUIC
	quicConf := &quic.Config{}
	//设置处理器
	handler := setupHandler("")

	var wg sync.WaitGroup
	wg.Add(len(bs))
	for _, b := range bs {
		bCap := b
		go func() {
			var err error
			//配置HTTP3
			server := http3.Server{
				Server:     &http.Server{Handler: handler, Addr: bCap},
				QuicConfig: quicConf,
			}
			err = server.ListenAndServeTLS(testdata.GetCertificatePaths())
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

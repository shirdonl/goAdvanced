package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"gitee.com/shirdonl/goAdvanced/chapter4/testdata"
	"github.com/lucas-clemente/quic-go"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	//解析命令行参数
	flag.Parse()
	urls := flag.Args()

	//HTTP3 密钥配置
	pool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	testdata.AddRootCA(pool)

	var qconf quic.Config
	//HTTP3配置
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
		},
		QuicConfig: &qconf,
	}
	defer roundTripper.Close()
	httpClient := &http.Client{
		Transport: roundTripper,
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, addr := range urls {
		fmt.Printf("GET %s", addr)
		go func(addr string) {
			rsp, err := httpClient.Get(addr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Got response for %s: %#v", addr, rsp)

			body := &bytes.Buffer{}
			_, err = io.Copy(body, rsp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Request Body:")
			fmt.Printf("%s", body.Bytes())
			wg.Done()
		}(addr)
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"gitee.com/shirdonl/goAdvanced/chapter2/advanced/thumb"
	"log"
	"os"
	"sync"
	"time"
)

func generateThumbnails(filenames []string) error {
	// errors 是一个无缓冲的 channel
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumb.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

func main() {
	now := time.Now()
	err := generateThumbnails(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("elapse:%.3f ms\n", 1000*time.Since(now).Seconds())
}

func generateThumbnail(filenames []string) {
	for _, f := range filenames {
		if _, err := thumb.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func generateThumbnail2(filenames []string) {
	for _, f := range filenames {
		go thumb.ImageFile(f)
	}
}

//并行生成指定文件的缩略图
func generateThumbnail3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		//启动goroutine
		go func(f string) {
			thumb.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	//等待goroutines完成
	for range filenames {
		<-ch
	}
}

func generateThumbnail4(filenames []string) {
	//声明等待组
	var wg sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		//启动goroutine
		go func(f string) {
			defer wg.Done()
			thumb.ImageFile(f)
		}(f)
	}
	//等待组等待
	wg.Wait()
}

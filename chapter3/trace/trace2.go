package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	// 使用 GOMAXPROCS设置可以同时执行的cpu的最大数量 为 1 个
	runtime.GOMAXPROCS(1)

	f, _ := os.Create("myTrace.out")
	defer f.Close()

	//开始跟踪，在跟踪时，跟踪将被缓冲并写入 一个我们指定的文件中
	_ = trace.Start(f)
	defer trace.Stop()

	// 咱们自定义一个任务
	ctx, task := trace.NewTask(context.Background(), "customerTask")
	defer task.End()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// 启动10个协程，模拟做任务
		go func(num string) {
			defer wg.Done()

			// 标记  num
			trace.WithRegion(ctx, num, func() {
				var sum, i int64
				// 模拟执行任务
				for ; i < 5000; i++ {
					sum += i
				}
				fmt.Println(num, sum)
			})
		}(fmt.Sprintf("num_%02d", i))
	}
	wg.Wait()
}

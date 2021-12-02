package bench

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func CgoBenchmark(b *testing.B) {
	CallCgo(b.N)
}

// 必须使用“-gcflags-l”调用GoBenchmark以避免内联
func GoBenchmark(b *testing.B) {
	CallGo(b.N)
}

// 演示了每个阻塞的cgo调用都消耗一个新线程
func Example_cgo_sleep() {
	const maxThreads = 50

	var wg sync.WaitGroup
	wg.Add(2 * maxThreads)
	for i := 0; i < 2*maxThreads; i++ {
		StartSleeper(&wg)
	}
	// 确保所有的goroutine都在运行
	wg.Wait()

	buf := make([]byte, 1<<16)
	buf = buf[:runtime.Stack(buf, true)]
	if bytes.Count(buf, []byte("locked to thread")) > maxThreads {
		fmt.Printf("> %d threads are running", maxThreads)
	}
}

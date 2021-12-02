package main

import (
	"fmt"
	"time"
)

type Request interface{}

func handle(r Request) { fmt.Println(r.(int)) }

const RateLimitPeriod = time.Minute

//任何一分钟内最多处理100个请求
const RateLimit = 100

//处理多个请求的函数
func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func main() {
	requests := make(chan Request)
	go handleRequests(requests)
	for i := 0; ; i++ {
		requests <- i
	}
}

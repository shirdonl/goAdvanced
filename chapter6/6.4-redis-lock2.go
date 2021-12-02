package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sync"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()
var mutex sync.Mutex
var counter int64
var wg sync.WaitGroup

func NewRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", //空密码
		DB:       0,  //默认的DB
	})
}
func Lock(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	bool, err := rdb.SetNX(ctx, key, 1, 10*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return bool
}
func UnLock(key string) int64 {
	nums, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			Lock("key")
		}()
	}
	wg.Wait()
	fmt.Printf("final counter is %d\n", counter)
}

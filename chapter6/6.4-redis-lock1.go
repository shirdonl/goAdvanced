package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       0,
})

var counter int64
var wg sync.WaitGroup
var rdb *redis.Client
var ctx = context.Background()

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			Lock(Incr)
		}()
	}
	wg.Wait()
}

func Incr() {
	counter++
}

func Lock(myfunc func()) {
	defer wg.Done()

	//上锁
	lockSuccess, err := client.SetNX(ctx, "test_key", 2, time.Second*5).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get unlock fail")
	} else {
		fmt.Println("get lock")
	}

	//运行
	myfunc()

	//解锁
	_, err = client.Del(ctx, "key").Result()
	if err != nil {
		fmt.Println(err)
		fmt.Println("unlock fail")
	} else {
		fmt.Println("unlock")
	}
}

func getUuid() string {
	return "TEST_KEY"
}
func lock1(myfunc func()) { //lock
	uuid := getUuid()
	lockSuccess, err := client.SetNX(ctx, "TEST_KEY", 6, time.Second*5).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail")
		return
	} else {
		fmt.Println("get lock")
	}
	myfunc()
	value, _ := client.Get(ctx, "TEST_KEY").Result()
	if value == uuid { //compare value,if equal then del
		_, err := client.Del(ctx, "TEST_KEY").Result()
		if err != nil {
			fmt.Println("unlock fail")
		} else {
			fmt.Println("unlock")
		}
	}
}

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

package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisMq struct {
	RedisClient *redis.Pool
	redisHost   string
	redisDB     int
	Mgo         *Mgo
	C           redis.Conn
}

func InitRedisMq(redisHost string, redisDB int) (*RedisMq, error) {
	rmq := &RedisMq{
		redisHost: redisHost,
		redisDB:   redisDB,
	}
	c, err := redis.Dial("tcp", redisHost)
	if err != nil {
		// handle error
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go InitRedisMq error: " + err.Error())
		return nil, err
	}
	c.Do("SELECT", rmq.redisDB)
	rmq.C = c
	rmq.Mgo = InitMgoDB("localhost:27017", "urls")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go InitRedisMq success ")
	return rmq, nil
}

/*
 * this function is likely a producter.
 */
func (rmq *RedisMq) LoadUrlsFromRedis(jobChan chan string) {
	//1) load Data
	//2) dispatchjob
	// When finish you need dispatchjob for
	// every blocked work because of none data in redis
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadUrlsFromRedis begin")
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go LoadUrlsFromRedis select")
			go rmq.GetUrls(jobChan)
			// urls := rmq.GetUrls(jobChan)
			// if len(urls) == 0 {
			//   fmt.Println("loadUrlsFromRedis urls is nil sleep 60s")
			//   time.Sleep(60 * time.Second)
			// }
			// fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadUrlsFromRedis add " + urls[1] + " to jobChan")
			// jobChan <- urls[1]
		}
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadUrlsFromRedis end")
}

func (rmq *RedisMq) GetUrls(jobChan chan string) {
	c := rmq.C
	// defer rc.Close()
	//values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
	n, _ := redis.Int(c.Do("LLEN", "url"))

	fmt.Println(time.Now().Format("2006-01-02 15:04:05")+" redismq.go GetUrls url num: %d", n)
	// if len(urls) < 100 then load data from mongodb
	if n < 100 {
		rmq.loadDataFromMongod(100 - n)
	}
	url, err := redis.String(c.Do("rpop", "url"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go GetUrls url: " + url)

	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go GetUrls url error: " + err.Error())
		return
	}
	jobChan <- url
	// return urls
}

func (rmq *RedisMq) loadDataFromMongod(n int) {
	//1) queru 1000 urls from mongodb
	//2) push urls to redismq
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadDataFromMongod begin")
	urls, _ := rmq.Mgo.QueryUrls(n)
	for _, url := range urls {
		rmq.PushUrl(url)
		rmq.Mgo.DeleteUrl(url)
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadDataFromMongod end")
}

func (rmq *RedisMq) PushUrl(url Url) {
	// rc := rmq.RedisClient.Get()
	rc := rmq.C
	// defer rc.Close()
	//values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
	rc.Do("lpush", "url", url.Url)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go PushUrl url: " + url.Url)
}

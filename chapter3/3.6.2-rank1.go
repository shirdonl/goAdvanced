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
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()

	//集合中增加用户分数
	_, err = conn.Do("ZADD", "score", 99, "Jack")
	if err != nil {
		panic(err)
	}
	//集合中批量增加用户分数
	_, err = conn.Do("ZADD", "score", 97, "James", 85, "Shirdon")
	if err != nil {
		panic(err)
	}

	//获取成员个数
	result, err := conn.Do("ZCARD", "score")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	//取出并升序
	scoreMap, err := redis.StringMap(conn.Do("ZREVRANGE", "score", 0, 2, "withscores"))
	for name := range scoreMap {
		fmt.Println(name, scoreMap[name])
	}

	//取出并降序
	scoreMap, err = redis.StringMap(conn.Do("ZRANGE", "score", 0, 1, "withscores"))
	for name := range scoreMap {
		fmt.Println(name, scoreMap[name])
	}

	//取出 Shirdon的分数
	score, err := redis.Int(conn.Do("ZSCORE", "score", "Shirdon"))
	if err != nil {
		panic(err)
	}
	fmt.Println(score)

	//移除集合中的某一个或者多个成员
	result, err = conn.Do("ZREM", "score", "Shirdon")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

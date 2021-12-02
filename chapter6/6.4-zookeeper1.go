package main

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	//上锁
	err = l.Lock()
	if err != nil {
		panic(err)
	}
	println("上锁成功, 中间输入逻辑处理语句")

	time.Sleep(time.Second * 10)

	//解锁
	l.Unlock()
	println("解锁成功, 完成逻辑处理语句")
}

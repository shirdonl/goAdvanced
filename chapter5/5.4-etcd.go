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
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
)

//Etcd互斥结构体
type EtcdMutex struct {
	Ttl     int64              //租约时间
	Conf    clientv3.Config    //etcd集群配置
	Key     string             //etcd的key
	cancel  context.CancelFunc //关闭续租的func
	lease   clientv3.Lease
	leaseID clientv3.LeaseID
	txn     clientv3.Txn
}

//初始化锁
func (em *EtcdMutex) init() error {
	var err error
	//定义上下午对象
	var ctx context.Context
	client, err := clientv3.New(em.Conf)
	if err != nil {
		return err
	}
	//创建事务
	em.txn = clientv3.NewKV(client).Txn(context.TODO())
	em.lease = clientv3.NewLease(client)
	//创建响应对象
	leaseResp, err := em.lease.Grant(context.TODO(), em.Ttl)
	if err != nil {
		return err
	}
	ctx, em.cancel = context.WithCancel(context.TODO())
	em.leaseID = leaseResp.ID
	_, err = em.lease.KeepAlive(ctx, em.leaseID)
	return err
}

//加锁
func (em *EtcdMutex) Lock() error {
	//初始化
	err := em.init()
	if err != nil {
		return err
	}
	//加锁
	em.txn.If(clientv3.Compare(clientv3.CreateRevision(em.Key), "=", 0)).
		Then(clientv3.OpPut(em.Key, "", clientv3.WithLease(em.leaseID))).
		Else()
	//提交
	_, err = em.txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

//释放锁
func (em *EtcdMutex) UnLock() {
	em.cancel()
	em.lease.Revoke(context.TODO(), em.leaseID)
	fmt.Println("释放了锁")
}

func main() {
	//定义终点
	var conf = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379", "127.0.0.1:2380"},
		DialTimeout: 5 * time.Second,
	}
	//初始化 eMutex1
	eMutex1 := &EtcdMutex{
		Conf: conf,
		Ttl:  10,
		Key:  "lock",
	}
	//初始化 eMutex2
	eMutex2 := &EtcdMutex{
		Conf: conf,
		Ttl:  10,
		Key:  "lock",
	}
	//启动 groutine1
	go func() {
		err := eMutex1.Lock()
		if err != nil {
			fmt.Println("groutine1 抢锁失败～")
			fmt.Println(err)
			return
		}
		fmt.Println("groutine1 抢锁成功～")
		time.Sleep(10 * time.Second)
		defer eMutex1.UnLock()
	}()

	//启动 groutine2
	go func() {
		err := eMutex2.Lock()
		if err != nil {
			fmt.Println("groutine2 抢锁失败～")
			fmt.Println(err)
			return
		}
		fmt.Println("groutine2 抢锁成功～")
		defer eMutex2.UnLock()
	}()
	time.Sleep(30 * time.Second)
}

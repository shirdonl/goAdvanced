package distribute

import (
	"fmt"
	"gitee.com/shirdonl/goAdvanced/chapter4/crawler"
	"net"
	"net/http"
	"net/rpc"
)

//定义子任务结构体
type Worker struct {
	addr          string
	addUrlChannel chan bool
}

//初始化子任务
func initWorker(addr string) *Worker {
	w := &Worker{}
	w.addr = addr
	w.addUrlChannel = make(chan bool)
	return w
}

//运行子任务
func RunWorker(mAddr, wAddr string) {
	fmt.Println("=======子任务开始=======")
	w := initWorker(wAddr)
	//开启RPC任务
	go startRpcWorker(w)
	register(mAddr, w.addr)
	fmt.Println("=======子任务结束=======")
	for {
		select {
		case <-w.addUrlChannel:
			fmt.Println("添加URL到通道成功")
		}
	}
}

//注册
func register(mAddr, wAddr string) {
	args := &RegisterArgs{}
	args.Worker = wAddr
	var reply RegisterReply
	call(mAddr, "Master.Register", args, &reply)
}

//执行RPC子任务
func (w *Worker) DoJob(args *DoJobArgs, res *DoJobReply) error {
	switch args.JobType {
	case "Crawler":
		//执行爬虫
		crawler.Do(args.Urls)
		//添加URL到通道成功
		w.addUrlChannel <- true
	}
	return nil
}

//开启RPC任务
func startRpcWorker(w *Worker) {
	//注册RPC
	rpc.Register(w)
	//处理HTTP
	rpc.HandleHTTP()
	err := http.ListenAndServe(w.addr, nil)
	fmt.Println("注册RPC错误：", err)
	rpcServer := rpc.NewServer()
	rpcServer.Register(w)
	_, e := net.Listen("tcp", w.addr)
	if e != nil {
		fmt.Println("运行子任务：", w.addr, " 错误： ", e)
	}
}

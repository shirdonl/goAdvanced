package distribute

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

//定义主任务结构体
type Master struct {
	addr         string
	regChan      chan string
	workDownChan chan string
	jobChan      chan string
	workers      map[*WorkInfo]bool
}

//定义子任务信息结构体
type WorkInfo struct {
	workAddr string
}

//初始化
func initMaster(addr string) (m *Master, err error) {
	m = &Master{}
	m.addr = addr
	m.regChan = make(chan string)
	m.jobChan = make(chan string, 2)
	m.workers = make(map[*WorkInfo]bool)
	return m, err
}

//获取URL
func getUrls(jobChan chan string) {
	for i := 0; i <= 10; i++ {
		url := "https://github.com/search?q=go&type=Repositories&p=1" + strconv.Itoa((i-1)*50)
		jobChan <- url
	}

}

//运行主服务器
func RunMaster(addr string) {
	m, err := initMaster(addr)
	if err != nil {
		fmt.Println("注册主任务错误：" + err.Error())
		return
	}

	//启动RPC主任务
	go startRpcMaster(m)
	getUrls(m.jobChan)

	//定义子任务地址
	workAddr := "127.0.0.1:8082"
	work := &WorkInfo{workAddr: workAddr}
	m.workers[work] = true
	fmt.Println("注册子任务：", work.workAddr)

	//分配子任务
	dispatchJob(work, m)

	fmt.Println("=======主任务运行结束=======")
	for {
		select {
		//当通道获取到子任务发送过来的地址
		case workAddr := <-m.regChan:
			work := &WorkInfo{workAddr: workAddr}
			m.workers[work] = true
			fmt.Println("注册子任务： ", work.workAddr)

			dispatchJob(work, m)
		}
	}
}

//分发任务
func dispatchJob(workInfo *WorkInfo, m *Master) {
	//组装RPC参数
	//urls := []string{"www.baidu.com"}

	var urls []string
	for i := 0; i < 10; i++ {
		url := <-m.jobChan // 从通道中获取URL
		urls = append(urls, url)
	}
	args := &DoJobArgs{}

	args.JobType = "Crawler"
	args.Urls = urls
	//定义任务回复结构体
	var reply DoJobReply
	//RPC调用子任务
	err := call(workInfo.workAddr, "Worker.DoJob", args, &reply)
	if err == true {
		m.workers[workInfo] = false
		fmt.Println("分配任务成功到地址: " + workInfo.workAddr)
	}
}

//注册RPC子任务
func (m *Master) Register(args *RegisterArgs, res *RegisterReply) error {
	m.regChan <- args.Worker
	return nil
}

func startRpcMaster(m *Master) {
	rpc.Register(m)
	rpc.HandleHTTP()
	err := http.ListenAndServe(m.addr, nil)
	if err != nil {
		fmt.Println("注册服务器错误：接收错误：", err)
	}
	rpcServer := rpc.NewServer()
	rpcServer.Register(m)
	_, e := net.Listen("tcp", m.addr)
	if e != nil {
		fmt.Println("运行子任务：", m.addr, " 错误： ", e)
	}
	fmt.Println("开启RPC主任务成功！")
}

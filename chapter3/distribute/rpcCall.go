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

package distribute

import (
	"fmt"
	"net/rpc"
)

//定义RPC参数
type DoJobArgs struct {
	JobType string
	Urls    []string
}

//定义RPC调用返回结果
type DoJobReply struct {
	OK bool
}

//定义RPC注册参数
type RegisterArgs struct {
	Worker string
}

//定义RPC注册返回参数
type RegisterReply struct {
	OK bool
}

//定义RPC调用方法
func call(srv string, rpcName string,
	args interface{}, reply interface{}) bool {
	c, err := rpc.DialHTTP("tcp", srv)
	if err != nil {
		return false
	}
	defer c.Close()

	err = c.Call(rpcName, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}

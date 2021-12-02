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

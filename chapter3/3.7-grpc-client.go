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
	pb "gitee.com/shirdonl/goAdvanced/chapter3/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8068", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}

	defer conn.Close()

	// 实例化
	client := pb.NewStudentServiceClient(conn)

	// 调用服务
	req := new(pb.Request)
	req.Name = "barry"
	resp, err := client.GetStudentInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}

	fmt.Printf("Recevied: %v\n", resp)
}

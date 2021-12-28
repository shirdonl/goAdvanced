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
	"net"
)

//定义服务结构体
type StudentServiceServer struct{}

func (p *StudentServiceServer) GetStudentInfo(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "barry" {
		resp = &pb.Response{
			Uid:      8,
			Username: name,
			Grade:    "6",
			GoodAt:   []string{"语文", "英语", "数学", "计算机"},
		}

	}
	err = nil
	return
}

func main() {
	port := ":8068"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()
	//将 StudentService 注册到 gRPC
	pb.RegisterStudentServiceServer(s, &StudentServiceServer{})
	s.Serve(l)
}

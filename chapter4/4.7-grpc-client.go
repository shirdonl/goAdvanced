package main

import (
	"fmt"
	pb "gitee.com/shirdonl/goAdvanced/chapter4/protobuf"
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

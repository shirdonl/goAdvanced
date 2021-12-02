package main

import (
	"fmt"
	"time"

	"gitee.com/shirdonl/goAdvanced/chapter6/micro/etcdv3/balancer"
	"gitee.com/shirdonl/goAdvanced/chapter6/micro/etcdv3/balancer/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	r := balancer.NewResolver("localhost:2378")
	resolver.Register(r)

	conn, err := grpc.Dial(r.Scheme()+"://author/project/test", grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewHelloServiceClient(conn)

	for {
		resp, err := client.Echo(context.Background(), &pb.Payload{Data: "hello"}, grpc.FailFast(true))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}

		<-time.After(time.Second)
	}
}

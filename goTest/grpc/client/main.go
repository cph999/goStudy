package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gotest/grpc/hello"
	_ "gotest/grpc/hello"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := hello.NewHelloClient(conn)

	// 调用方法
	req := &hello.HelloRequest{Name: "gRPC巴嘎"}
	res, err := c.SayHello(context.Background(), req)
	fmt.Println(res)
	if err != nil {
		grpclog.Fatalln(err)
	}
}

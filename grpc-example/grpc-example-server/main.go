// description:
// @author renshiwei
// Date: 2022/9/30 17:54

package main

import (
	helloGrpc "github.com/qiaoshurui/go-learn/grpc-example/grpc-example-proto/pb/hello"
	"github.com/qiaoshurui/go-learn/grpc-example/grpc-example-server/app/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	initGrpcServer()
}

func initGrpcServer() {
	grpcPort := ":8881"
	server := grpc.NewServer()

	// 注册服务类
	helloGrpc.RegisterHelloServer(server, hello.NewHelloService())

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Port port listen err:%+v", err)
		return
	}

	log.Printf("grpc server init... port%s", grpcPort)

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Port server init err:%+v", err)
		return
	}
}

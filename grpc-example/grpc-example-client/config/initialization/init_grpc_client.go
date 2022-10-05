// description:
// @author renshiwei
// Date: 2022/9/30 19:00

package initialization

import (
	"fmt"
	"github.com/qiaoshurui/go-learn/grpc-example/grpc-example-client/common/grpc"
	"github.com/qiaoshurui/go-learn/grpc-example/grpc-example-client/config"
	helloGrpc "grpc-example-proto/pb/hello"
	"log"
	"time"
)

func InitHelloGrpcClient() {
	grpcPort := ":8881"
	grpcClient, err := grpc.InitGrpcClient(&grpc.GrpcClientConfig{
		Host:           "127.0.0.1",
		Port:           grpcPort,
		ConnectTimeout: 60 * time.Second,
		RequestTimeout: 30 * time.Second,
	})
	if err != nil {
		panic(fmt.Sprintf("Grpc client init panic:%+v.", err))
	}

	log.Printf("hello grpc client init... port%s", grpcPort)

	// 注册 hello grpc
	config.HelloGrpcClient = helloGrpc.NewHelloClient(grpcClient.Conn)
}

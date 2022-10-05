// description:
// @author renshiwei
// Date: 2022/9/30 18:08

package hello

import (
	"context"
	"fmt"
	helloGrpc "grpc-example-proto/pb/hello"
	"log"
)

type Hello interface {
	SayHello(ctx context.Context, in *helloGrpc.HelloRequest) (*helloGrpc.HelloResponse, error)
}

type HelloAny struct{}

func NewHelloService() Hello {
	return &HelloAny{}
}

func (h *HelloAny) SayHello(ctx context.Context, helloRequest *helloGrpc.HelloRequest) (*helloGrpc.HelloResponse, error) {
	data := fmt.Sprintf("Hello %s!", helloRequest.Name)

	log.Printf("hello server grpc request param:%s.", helloRequest.Name)

	return &helloGrpc.HelloResponse{
		Code: "I2000",
		Msg:  "Success",
		Data: data,
	}, nil
}

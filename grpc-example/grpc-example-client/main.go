// description:
// @author renshiwei
// Date: 2022/9/30 17:54

package main

import (
	"github.com/qiaoshurui/go-learn/grpc-example/grpc-example-client/app/hello"
	"github.com/qiaoshurui/go-learn/grpc-example/grpc-example-client/config/initialization"
)

func main() {
	initialization.InitHelloGrpcClient()

	hello.SayHello()
}

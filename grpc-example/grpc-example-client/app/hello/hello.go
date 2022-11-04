// description:
// @author renshiwei
// Date: 2022/9/30 18:08

package hello

import (
	"bufio"
	"context"
	"fmt"
	"github.com/qiaoshurui/go-learn/grpc-example/grpc-example-client/config"
	helloGrpc "grpc-example-proto/pb/hello"
	"log"
	"os"
)

func SayHello() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name := scanner.Text()
		helloResponse, err := config.HelloGrpcClient.SayHello(context.Background(), &helloGrpc.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("grpc request err:%+v\n", err)
		}
		if helloResponse.Code == "I2000" {
			log.Printf("Hello Server response:%+v\n", helloResponse)
			fmt.Println(helloResponse.Data)
		} else {
			log.Printf("Hello Server response fail:%+v\n", helloResponse)
		}
	}
}

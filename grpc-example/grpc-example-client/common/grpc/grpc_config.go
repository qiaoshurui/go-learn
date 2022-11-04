// description:
// @author renshiwei
// Date: 2022/9/30 18:45

package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type GrpcClient struct {
	Config *GrpcClientConfig
	Conn   *grpc.ClientConn
}

type GrpcClientConfig struct {
	Host           string        `yaml:"host"`
	Port           string        `yaml:"grpc"`
	ConnectTimeout time.Duration `yaml:"connect_timeout"` // 链接超时时间（单位：秒）
	RequestTimeout time.Duration `yaml:"request_timeout"` // 请求超时时间（单位：秒）
}

func InitGrpcClient(cc *GrpcClientConfig) (*GrpcClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cc.ConnectTimeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, cc.Host+cc.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		// lb策略轮询模式
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		log.Fatalf("Grpc client init err:%+v.", err)
		return nil, err
	}
	return &GrpcClient{
		Conn:   conn,
		Config: cc,
	}, nil
}

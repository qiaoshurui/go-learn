module github.com/qiaoshurui/go-learn/grpc-example/grpc-example-server

go 1.18

replace grpc-example-proto v0.0.1 => github.com/qiaoshurui/go-learn/grpc-example/grpc-example-proto v0.0.0-20220930092846-bee87a58b594

require (
	google.golang.org/grpc v1.49.0
	grpc-example-proto v0.0.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

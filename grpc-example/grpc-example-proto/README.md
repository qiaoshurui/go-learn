
# 生成proto代码
在项目根目录 `go-leran` 下执行如下代码

## hello服务
```sh
protoc --go_out=./ -I=grpc-example/grpc-example-proto/proto/hello hello.proto

protoc --go-grpc_out=require_unimplemented_servers=false:./ -I=grpc-example/grpc-example-proto/proto/hello hello.proto
```

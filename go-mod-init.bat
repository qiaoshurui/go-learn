
cd ./grpc-example/grpc-example-proto

go mod tidy -compat=1.18

cd ../grpc-example-server

go mod tidy -compat=1.18

cd ../grpc-example-client

go mod tidy -compat=1.18

cd ..

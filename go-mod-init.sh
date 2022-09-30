#!/bin/bash

cd ./grpc-example/grpc-example-proto || exit

go mod tidy -compat=1.18

cd ../grpc-example-server || exit

go mod tidy -compat=1.18

cd ../grpc-example-client || exit

go mod tidy -compat=1.18

cd ..

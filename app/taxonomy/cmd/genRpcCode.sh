#!/bin/bash
goctl rpc protoc rpc/pb/$1.proto --go_out=rpc/pb --go-grpc_out=rpc/pb --zrpc_out=rpc

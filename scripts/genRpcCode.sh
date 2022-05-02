#!/bin/bash
goctl rpc protoc $1.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

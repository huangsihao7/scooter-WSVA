@echo off

goctl rpc protoc recommend.proto --go_out=. --go-grpc_out=. --zrpc_out=.
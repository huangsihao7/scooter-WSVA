@echo off

goctl rpc protoc feed.proto --go_out=. --go-grpc_out=. --zrpc_out=.
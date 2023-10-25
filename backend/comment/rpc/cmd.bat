@echo off

goctl rpc protoc comment.proto --go_out=. --go-grpc_out=. --zrpc_out=.
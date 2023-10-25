@echo off

goctl rpc protoc favorite.proto --go_out=. --go-grpc_out=. --zrpc_out=.
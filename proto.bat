@echo off
protoc -I ./third_party --proto_path=./apply/proto --go_out=./apply/proto --go_opt paths=source_relative --go-grpc_out=./apply/proto --go-grpc_opt paths=source_relative --grpc-gateway_out=./apply/proto --grpc-gateway_opt paths=source_relative  ./apply/proto/*.proto

:: protoc -I ./third_party --proto_path=./user/proto --go_out=./user/proto --go_opt paths=source_relative --go-grpc_out=./user/proto --go-grpc_opt paths=source_relative --grpc-gateway_out=./user/proto --grpc-gateway_opt paths=source_relative ./user/proto/*.proto

pause
## gen to folder server
## Opt : go_out for message
## Opt : grpc_out for service
 protoc calculator.proto --go_out=../server --go-grpc_out=../server
 protoc example.proto --go_out=../server --go-grpc_out=../server

## Build All
 protoc *.proto --go_out=../server --go-grpc_out=../server

## client
  protoc calculator.proto --go_out=../client --go-grpc_out=../client

  protoc example.proto --go_out=../client --go-grpc_out=../client

## Go get package in project
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Install gRPC tool in project

go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# 1. Install go package
 go get google.golang.org/protobuf       
 go get google.golang.org/grpc
 


 ## run gRPC client
evans --proto=./calculator.proto --host=localhost -p 8009
evans --proto=./example.proto --host=localhost -p 8009

## run eVan with Ca 
evans --proto=./calculator.proto --tls --cacert ca.crt --host=localhost -p 8009 
## not use 

# command
show services 
call fn

## ctr + d for stop send stram

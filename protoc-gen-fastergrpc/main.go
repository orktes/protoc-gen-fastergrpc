package main

import (
	"github.com/gogo/protobuf/vanity/command"
	grpc "github.com/orktes/protoc-gen-fastergrpc/grpc"
)

func main() {
	req := command.Read()
	resp := command.GeneratePlugin(req, grpc.NewGRPC(), "_fastgrpc.go")
	command.Write(resp)
}

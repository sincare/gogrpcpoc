package services

import (
	context "context"
	"fmt"
)

type foobarServer struct {
}

func NewFoobarServer() FoobarServer {
	return foobarServer{}
}

func (foobarServer) mustEmbedUnimplementedFoobarServer() {}

func (foobarServer) FooApi(ctx context.Context, req *FooRequest) (*FooResponse, error) {
	result := fmt.Sprintf("Result of foo, called by %v and age %v", req.Fooname, req.Fooage)

	return &FooResponse{
		Result: result,
	}, nil
}

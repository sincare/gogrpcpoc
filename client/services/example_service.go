package services

import (
	"context"
	"fmt"
)

type FoobarService interface {
	FooApi(fooname string, fooage int32) error
}

type foobarService struct {
	foobarClient FoobarClient
}

func NewFoobarService(foobarClient FoobarClient) FoobarService {
	return foobarService{foobarClient}
}

func (foo foobarService) FooApi(fooname string, fooage int32) error {
	req := FooRequest{
		Fooname: fooname,
		Fooage:  fooage,
	}

	resp, err := foo.foobarClient.FooApi(context.Background(), &req)

	if err != nil {
		return err
	}

	fmt.Println("==FooApi==")
	fmt.Printf("Request:%v %v \n", fooname, fooname)
	fmt.Printf("Response:%v", resp.Result)
	
	return nil
}

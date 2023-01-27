package services

import (
	"context"
	"fmt"
	"io"
	"time"
)

type CalculatorService interface {
	Hello(msg string) error
	Fibonacci(n uint32) error
	Average(nums ...float64) error
	Sum(nums ...int32) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient}
}

func (cs calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}

	resp, err := cs.calculatorClient.Hello(context.Background(), &req)

	if err != nil {
		return err
	}

	fmt.Println("Hellp", resp)

	return nil
}

func (cs calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	stream, err := cs.calculatorClient.Fibonacci(ctx, &req)

	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println("fib resp", res.Result)
	}

	return nil
}

func (cs calculatorService) Average(nums ...float64) error {
	stream, err := cs.calculatorClient.Average(context.Background())

	if err != nil {
		return err
	}

	for _, num := range nums {
		req := AverageRequest{
			Num: num,
		}

		stream.Send(&req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		return err
	}

	fmt.Println("Average", res)

	return nil
}

func (cs calculatorService) Sum(nums ...int32) error {
	stream, err := cs.calculatorClient.Sum(context.Background())

	if err != nil {
		return err
	}

	go func() {
		for _, num := range nums {
			req := SumRequest{
				Num: num,
			}

			stream.Send(&req)
		}

		stream.CloseSend()
	}()

	done := make(chan bool)
	errs := make(chan error)

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				errs <- err
			}
			fmt.Println("Response from stram", res.Result)
		}

		done <- true
	}()

	select {
	case <-done:// Meaning is true = <- done ?
		return nil
	case err := <-errs:
		return err
	}
}

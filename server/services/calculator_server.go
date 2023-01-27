package services

import (
	context "context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (server calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Name is required",
		)
	}
	msg := fmt.Sprintf("Hello man!! %v", req.Name)

	resp := HelloResponse{
		Result: msg,
	}

	//return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
	return &resp, nil
}

func (server calculatorServer) Fibonacci(req *FibonacciRequest, stream Calculator_FibonacciServer) error {
	for n := uint32(0); n < req.N; n++ {
		fib := fib(n)
		res := FibonacciResponse{
			Result: fib,
		}

		time.Sleep(time.Second)

		stream.Send(&res)
	}

	return nil
}

func fib(n uint32) uint32 {
	if n == 0 || n == 1 {
		return n
	}

	return (fib(n-1) + fib(n-2))
}

func (calculatorServer) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		sum += req.Num
		count++
	}

	res := AverageResponse{
		Result: sum / count,
	}

	return stream.SendAndClose(&res)
}

func (calculatorServer) Sum(stream Calculator_SumServer) error {
	var sum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF { // wait until end of req
			break
		}

		sum += req.Num

		// resp now
		res := SumResponse{
			Result: sum,
		}

		err = stream.Send(&res)

		if err == io.EOF {
			break
		}
	}

	return nil
}

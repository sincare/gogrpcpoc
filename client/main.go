package main

import (
	"client/services"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	var cc *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	host := flag.String("host", "localhost:8009", "gRPC Server host.")
	tls := flag.Bool("tls", false, "use a secure TLS Connection.")
	flag.Parse()

	if *tls {
		fmt.Println("Prepare TLS")
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatal(err)
		}

	} else {
		creds = insecure.NewCredentials() // on prod chose be credential
	}

	cc, err = grpc.Dial(*host, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()
	/*
			creds := insecure.NewCredentials() // on prod chose be credential
			cc, err := grpc.Dial("127.0.0.1:8009", grpc.WithTransportCredentials(creds))
			if err != nil {
			log.Fatal(err)
		}

	*/

	calClient := services.NewCalculatorClient(cc)
	calService := services.NewCalculatorService(calClient)

	foobarClient := services.NewFoobarClient(cc)
	forbarService := services.NewFoobarService(foobarClient)

	fmt.Println("Call forbarService")
	err = forbarService.FooApi("Ya A Da", 28)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Call Service [Unary]Hello")
	err = calService.Hello("MrMozza")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Call Service [Server Streaming]Fibonacci")
	err = calService.Fibonacci(5)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Call Service [Client Streaming]Average")
	err = calService.Average(1, 2, 3)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Call Service [Client Streaming]Average")
	//req := []float64{1, 2, 3, 4, 5}
	err = calService.Average(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Call Service [Bi-drection Streaming]Average")
	err = calService.Sum(1, 2, 3, 4, 5, 6, 7)

	if err != nil {
		log.Fatal(err)
	}

	// test handler err

	fmt.Println("Call Service [Unary]Hello")
	err = calService.Hello("")

	if err != nil {
		// It mean if Error from gRPC if yes, variable ok will be true
		if grpcErr, ok := status.FromError(err); ok {
			// if error from grpc  err can get Code and Message from error
			log.Printf("[%v] %v", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)

		}
	}

}

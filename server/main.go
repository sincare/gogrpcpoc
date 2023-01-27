package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	var sevr *grpc.Server
	//conf
	tls := flag.Bool("tls", false, "use a secure TLS connection.")
	flag.Parse()

	if *tls {
		fmt.Println("Prepare TLS")
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatal(err)
		}

		sevr = grpc.NewServer(grpc.Creds(creds))

	} else {
		sevr = grpc.NewServer()
	}

	listener, err := net.Listen("tcp", ":8009")

	if err != nil {
		log.Fatal(err)
	}

	// register for listener
	services.RegisterCalculatorServer(sevr, services.NewCalculatorServer())
	services.RegisterFoobarServer(sevr, services.NewFoobarServer())

	fmt.Print("\ngRPC server listening on port 8009")

	if *tls {
		fmt.Println(" with TLS.")
	}

	err = sevr.Serve(listener)
	// for client remember
	reflection.Register(sevr)

	if err != nil {
		log.Fatal(err)
	}
}

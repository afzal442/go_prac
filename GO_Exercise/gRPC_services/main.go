package main

import (
	"net"
	"os"

	"gRPC/bookInfo_server"
	pb "gRPC/protos/bookInfo"
	"gRPC/rates"

	"log"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {

	l := log.Default()

	rate, err := rates.NewRate(l)
	if err != nil {
		log.Printf("Unable to generate rates", "error", err)
		os.Exit(1)
	}

	// Create a new server
	gs := grpc.NewServer()

	bs := bookInfo_server.NewBookInfo(rate, l)

	// Register BookInfo service with the server
	pb.RegisterBookInfoServer(gs, bs)

	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// listen for requests
	gs.Serve(lis)

}

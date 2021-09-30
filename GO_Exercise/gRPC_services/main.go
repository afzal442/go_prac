package main

import (
	"net"

	"gRPC/bookInfo_server"
	pb "gRPC/protos/bookInfo"

	"log"

	"google.golang.org/grpc"
)

func main() {

	l := log.Default()

	// Create a new server
	gs := grpc.NewServer()

	bs := bookInfo_server.NewBookInfo(l)

	// Register the server
	pb.RegisterBookInfoServer(gs, bs)

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// listen for requests
	gs.Serve(lis)

}

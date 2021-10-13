package bookInfo_server

import (
	"context"
	"io"
	"log"
	"time"

	pb "gRPC/protos/bookInfo"
)

type BookInfo struct {
	l *log.Logger
}

// type BookInfo struct {
// 	pb.UnimplementedBookInfoServer
// }

func NewBookInfo(l *log.Logger) *BookInfo {
	return &BookInfo{l}
}

func (b *BookInfo) GetRate(ctx context.Context, xr *pb.RateRequest) (*pb.RateResponse, error) {
	// b.log.Info("Handle request for GetRate", "Article Name", xr.GetArticleName(), "Article Review", xr.GetArticleReview())
	log.Printf("Received: ", "Article Name", xr.GetTitle(), "Article Review", xr.GetReview())

	// return a response
	return &pb.RateResponse{Rating: 3}, nil

}

func (b *BookInfo) SubscribeRate(src pb.BookInfo_SubscribeRateServer) error {

	// handle client messages
	// Recv is a blocking method which returns on client data
	// io.EOF signals that the client has closed the connection

	// var er error
	// i := 0

	rx, _ := src.Recv()
	review := rx.GetReview()

	go func() {
		for {
			rx, err := src.Recv()

			if err == io.EOF {
				log.Println("Client has closed connection")
				break
			}

			// any other error means the transport between the server and client is unavailable
			if err != nil {
				log.Printf("Unable to read from client", "error", err)
				break
			}

			log.Printf("Handle client request", "request_base", rx.GetTitle(), "request_dest", rx.GetReview())
			review = rx.GetReview()
		}
	}()
	// handle server responses
	// we block here to keep the connection open
	// send a message back to the client

	// time.Sleep(time.Second * 3)

	for {
		switch review {
		case "Good":
			er := src.Send(&pb.RateResponse{Rating: 2})
			if er != nil {
				log.Printf("Unable to send to client", "error", er)
			}
			// break
		case "Very Good":
			er := src.Send(&pb.RateResponse{Rating: 3})
			if er != nil {
				log.Printf("Unable to send to client", "error", er)
			}
			// break
		case "Excellent":
			er := src.Send(&pb.RateResponse{Rating: 4})
			if er != nil {
				log.Printf("Unable to send to client", "error", er)
			}
			// break
		}
		time.Sleep(time.Second * 15)
	}
	return nil
}

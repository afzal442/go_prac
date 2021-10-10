package bookInfo_server

import (
	"context"
	"gRPC/rates"
	"log"

	pb "gRPC/protos/bookInfo"
)

type BookInfo struct {
	rating *rates.Rate
	l      *log.Logger
}

// type BookInfo struct {
// 	pb.UnimplementedBookInfoServer
// }

func NewBookInfo(m *rates.Rate, l *log.Logger) *BookInfo {
	return &BookInfo{m, l}
}

func (b *BookInfo) GetRate(ctx context.Context, xr *pb.RateRequest) (*pb.RateResponse, error) {
	// b.log.Info("Handle request for GetRate", "Article Name", xr.GetArticleName(), "Article Review", xr.GetArticleReview())
	log.Printf("Received: ", "Article Name", xr.GetTitle(), "Article Review", xr.GetReview())

	rt, err := b.rating.GetRatings(xr.GetTitle(), xr.GetReview())
	if err != nil {
		return nil, err
	}

	return &pb.RateResponse{Rating: rt}, nil

}

// func (b *BookInfo) SubscribeRate(src pb.BookInfo_SubscribeRateServer) error {

// 	// handle client messages
// 	rr, err := src.Recv() // Recv is a blocking method which returns on client data
// 	// io.EOF signals that the client has closed the connection

// 	if err == io.EOF {
// 		log.Printf("Client has closed connection")
// 	}

// 	// any other error means the transport between the server and client is unavailable
// 	if err != nil {
// 		log.Printf("Unable to read from client", "error", err)
// 	}

// 	log.Printf("Handle client request", "request_base", rr.GetArticleName(), "request_dest", rr.GetArticleReview())
// 	// handle server responses
// 	// we block here to keep the connection open
// 	// send a message back to the client

// 	time.Sleep(time.Second * 3)
// 	er := src.Send(&pb.RateResponse{Rating: 4})
// 	if er != nil {
// 		log.Printf("Unable to send", "error", er)
// 	}

// 	return nil

// }

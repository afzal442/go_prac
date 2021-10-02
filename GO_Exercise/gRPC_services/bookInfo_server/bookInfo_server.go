package bookInfo_server

import (
	"context"
	"log"

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
	log.Printf("Received: ", "Article Name", xr.GetArticleName(), "Article Review", xr.GetArticleReview())

	return &pb.RateResponse{Ratings: 4.5}, nil

}

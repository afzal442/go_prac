package rates

import (
	"encoding/json"
	"fmt"

	// pb "gRPC/protos/bookInfo"
	"log"
	"net/http"
)

type Rate struct {
	log     *log.Logger
	ratings map[string]string
}

func NewRate(l *log.Logger) (*Rate, error) {
	df := &Rate{
		log: l,
		ratings: map[string]string{
			"Good":      "3.5",
			"Very Good": "4.0",
			"Excellent": "5.0",
		},
	}
	err := df.GetRatings()

	return df, err
}

// func (r *Rate) GetRatings() error {
// 	resp, err := http.DefaultClient.Get("http://localhost:8080/")
// 	if err != nil {
// 		return err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("error getting ratings: %v", resp.StatusCode)
// 	}
// 	defer resp.Body.Close()

// 	md := &df{}

// 	json.NewDecoder(resp.Body).Decode(&md)

// 	for {
// 		dt := md.Rating

// 		r.ratings[md.Review] = dt

// 	}

// }

// func (r *Rate) GetRating(review string) (float64, error) {
// 	if r.ratings[review] == 0 {
// 		return 0, fmt.Errorf("no rating for %s", review)
// 	}
// 	return r.ratings[review], nil
// }

// func (r *Rate) AddRating(review string, rating float64) error {
// 	if r.ratings[review] != 0 {
// 		return fmt.Errorf("rating for %s already exists", review)
// 	}
// 	r.ratings[review] = rating
// 	return nil
// }

// func (r *Rate) UpdateRating(review string, rating float64) error {
// 	if r.ratings[review] == 0 {
// 		return fmt.Errorf("no rating for %s", review)
// 	}
// 	r.ratings[review] = rating
// 	return nil
// }

// func (r *Rate) DeleteRating(review string) error {
// 	if r.ratings[review] == 0 {
// 		return fmt.Errorf("no rating for %s", review)
// 	}
// 	r.ratings[review] = 0
// 	return nil
// }

type df struct {
	Review string `json:"review"`
	Rating string `json:"rating"`
}

func (r *Rate) GetRatings() error {
	resp, err := http.DefaultClient.Get("http://localhost:5000/articles")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error getting ratings: %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	md := &df{}

	json.NewDecoder(resp.Body).Decode(&md)

	for {
		dt := md.Rating

		r.ratings[md.Review] = dt

	}

}

// type Data struct{
// 	Dt []df ``
// }

// func (r *Rate) GetRating(ctx context.Context, req *pb.GetRatingRequest) (*pb.GetRatingResponse, error) {
// 	r.log.Printf("Received request for rating %s", req.Name)
// 	time.Sleep(1 * time.Second)
// 	rating, ok := r.ratings[req.Name]
// 	if !ok {
// 		return nil, fmt.Errorf("rating %s not found", req.Name)
// 	}
// 	return &pb.GetRatingResponse{
// 		Rating: rating,
// 	}, nil
// }

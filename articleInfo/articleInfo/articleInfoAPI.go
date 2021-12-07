package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type nwarticle struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Review string  `json:"review"`
	Rating float64 `json:"rating"`
}

// article represents data about a record article.
type article struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Review string  `json:"review"`
	Rating float64 `json:"rating"`
}

// articles slice to seed record article data.
var articles = []article{
	{ID: "1", Title: "Let's GO", Author: "John Doe", Price: 56.99, Review: "Very Good", Rating: 4},
	{ID: "2", Title: "the three mistakes of my life", Author: "Chetan", Price: 17.99, Review: "Good", Rating: 3.5},
	{ID: "3", Title: "3 idiots", Author: "Bhagat", Price: 39.99, Review: "Excellent", Rating: 5},
}

var iRegexp = regexp.MustCompile(`[1-9]`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	i := request.QueryStringParameters["list"]
	x, _ := strconv.Atoi(i)
	if iRegexp.MatchString(i) {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintln(articles[x-1]),
			StatusCode: 200,
		}, nil
	} else {
		return clientError(http.StatusBadRequest)
	}

	// v := request.PathParameters["id"]
	// if v == "viewAll" {
	// 	return list(request)
	// }

	// if request.Path == "/articleInfo/viewAll" {
	// 	return list(request)
	// }

	// if request.HTTPMethod == "GET" {
	// 	return list(request)
	// } else if request.HTTPMethod == "POST" {
	// 	return create(request)
	// }

	// how to handle request paths

	// if request.Path != "/articleInfo/view" {
	// 	return events.APIGatewayProxyResponse{}, errors.New("Non 200 Response found")
	// } else {
	// 	return events.APIGatewayProxyResponse{
	// 		Body:       fmt.Sprintf("%v", articles),
	// 		StatusCode: 200,
	// 	}, nil
	// }

	// if resp.StatusCode != 200 {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// ip, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// if len(ip) == 0 {
	// 	return events.APIGatewayProxyResponse{}, ErrNoIP
	// }

	// var newarticle article

	// reqBody, _ := ioutil.ReadAll(request.Body)
	// json.Unmarshal(reqBody, &newarticle)

	// articles = append(articles, newarticle)

}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// Similarly add a helper for send responses relating to client errors.
func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

// func list(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	return events.APIGatewayProxyResponse{
// 		Body:       fmt.Sprintf("%v", articles),
// 		StatusCode: 200,
// 	}, nil
// }

// func create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	var newarticle article

// 	var req *http.Request
// 	reqBody, _ := ioutil.ReadAll(req.Body)
// 	json.Unmarshal(reqBody, &newarticle)

// 	articles = append(articles, newarticle)

// 	return events.APIGatewayProxyResponse{
// 		Body:       fmt.Sprintf("%v", articles),
// 		StatusCode: 200,
// 	}, nil
// }

func main() {
	lambda.Start(handler)
}

// b := make([]byte, 32)

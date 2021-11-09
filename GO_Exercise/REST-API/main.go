package main

import (
	"REST-API/routes"
)

func main() {

	// conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	// if err != nil {
	// 	panic(err)
	// }

	// // create client
	// cc := pb.NewCurrencyClient(conn)

	// // create database instance
	// db := data.NewProductsDB(cc, l)

	// // create the handlers
	// ph := handlers.NewProducts(l, v, db)

	// defer conn.Close()

	// Initialize the routes
	r := routes.InitializeRoutes()

	// Start serving the application
	r.Run()
}

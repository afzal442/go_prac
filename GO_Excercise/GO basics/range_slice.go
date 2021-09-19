package main

import (
	"fmt"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

func main() {
	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	fmt.Println("Albums:")
	for _, album := range albums {
		fmt.Println(album)
	}
}

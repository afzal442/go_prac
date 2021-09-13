package main

import "fmt"

func main() {
	m := make(chan int, 2)
	c := make(chan string, 2)
	c <- "hey"
	m <- 1778
	//v := <-m
	fmt.Println(<-c)
	fmt.Println(<-m)
}

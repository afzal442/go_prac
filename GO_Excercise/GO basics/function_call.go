package main

import "fmt"

func cars(x string, y int) (string, int) {
	a := x
	b := y
	
	return a, b
}

func main() {
	fmt.Println(cars("hi", 5))
}
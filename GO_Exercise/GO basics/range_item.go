package main

import "fmt"

func main() {
	x := make([]int, 5) // Dynamic Array declaration
	for i := range x {
		x[i] = 3 << uint(i) // == 3*i
	}
	for _, value := range x {
		fmt.Printf("%d\t", value)
	}
}

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(i int) int {
	// x := 0
	var y [10]int
	return func(i int) int {
		if i == 0 {
			y[i] = 0
			return 0
		} else if i == 1 {
			y[i] = 1
			return 1
		} else {
			y[i] = y[i-1] + y[i-2]
			return y[i]
		}
		//return y[i]
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

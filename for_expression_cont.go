package main

import "fmt"

func check_sum(s int) int {
	for s < 1000 {
		s += s
	}
	return s
}

func main() {
	s := 1
	fmt.Println(check_sum(s))
}

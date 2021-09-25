package main

import "fmt"

func main() {
	var s []int
	s = []int{6, 7}
	fmt.Println(s) //

	s = append(s, 0)
	fmt.Println(s)

	// The slice grows as needed.
	s = append(s, 5)
	fmt.Println(s)

	// We can add more than one element at a time.
	s = append(s, 1, 4)
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v \t", s[i])
	}
}

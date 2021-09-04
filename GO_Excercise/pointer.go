package main

import "fmt"

func main() {
	i, j := 5, 10

	p := &i
	*p = *p + 6

	fmt.Println(i, *p) // 11, 11
	i = i + 2
	fmt.Println(i, *p) //13, 13

	p = &j
	*p = *p * 2
	j = j + 10
	fmt.Println(j) // 30
}

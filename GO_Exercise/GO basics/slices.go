package main

import "fmt"

func main() {
	kloud1_dept := [5]string{
		"Finance",
		"Analyst",
		"Project",
		"Engineering",
		"Research",
	}
	fmt.Println(kloud1_dept) // [Finance Analyst Project Engineering Research]

	x := kloud1_dept[:2]
	y := kloud1_dept[1:3]
	z := kloud1_dept[2:]
	fmt.Println(x, y, z) // [Finance Analyst] [Project Engineering] [Research]

	y[0] = "Business"
	fmt.Printf("%v %v", x, y)
	fmt.Println(kloud1_dept) // [Finance Business Project Engineering Research]
}

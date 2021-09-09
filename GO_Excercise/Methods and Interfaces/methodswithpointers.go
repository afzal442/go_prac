package main

import (
	"fmt"
	"math"
)

type kloud struct {
	x, y float64
}

func find_power(ac kloud) float64 {
	return math.Pow(ac.x, 2) + math.Pow(ac.y, 2)
}

func pointer_fun(ac *kloud, f float64) {
	ac.x = ac.x + f
	ac.y = ac.y + f
}

func main() {
	ac := kloud{1, 3}
	pointer_fun(&ac, 10)
	fmt.Println(find_power(ac))
}

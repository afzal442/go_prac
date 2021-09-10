package main

import (
	"fmt"
	"math"
)

type check_max struct {
	x, y float64
}

func max(v check_max) float64 {
	return math.Max(v.x, v.y)
}

func main() {
	v := check_max{3, 4}
	fmt.Println(max(v))
}

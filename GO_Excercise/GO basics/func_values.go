package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	sum := func(x, y float64) float64 {
		return math.Min(x, y)
	}
	fmt.Println(sum(5, 12))

	fmt.Println(compute(sum))
	fmt.Println(compute(math.Pow))
}

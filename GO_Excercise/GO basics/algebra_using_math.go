package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4                           // x = 3, y = 4
	var m float64 = math.Sqrt(float64(x*x + y*y)) // Pythagorean theorem
	var n uint32 = uint(m)                        //type conversion
	fmt.Println(x, y, n)                          // 3 4 5
}

package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy) // dynamic 2-D array dec with given length
	//fmt.Println((ss[0][3]))
	for y := 0; y < dy; y++ {
		//fmt.Println(y)
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			//fmt.Println(ss[y][x])
			s[x] = uint8((x + y) / 2) // convert to uint8
		}
		ss[y] = s // allocate 1D array to 2D array
	}
	fmt.Println(ss)
	return ss

}

func main() {
	pic.Show(Pic)
}

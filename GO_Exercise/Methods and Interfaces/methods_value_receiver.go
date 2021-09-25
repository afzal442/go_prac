package main

import (
	"fmt"
)

type kloud1 struct {
	x, y uint32
}

func (m *kloud1) find_mul(t uint32) {
	m.x = m.x * t
	m.y = m.y * t
}

func (m *kloud1) find_sum() uint32 {
	return (m.x*m.x + m.y*m.y)
}

func main() {
	m := &kloud1{3, 4}
	fmt.Printf("Before scaling: %+v, find_sum: %v\n", m, m.find_sum()) // 25
	(*m).find_mul(5)                                                   // or m.find_mul(5)
	fmt.Printf("After scaling: %+v, find_sum: %v\n", m, m.find_sum())  // 625
}

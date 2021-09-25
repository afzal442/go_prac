package main

import "fmt"

type I interface {
	M()
}

type hog struct {
	S string
	t int
}

// func (t *hog) M() {
func (t hog) M() {
	fmt.Println(t.S, t.t) // Prints "Gill 129"
}

func main() {
	var i I
	i = &hog{"Gill", 129}
	i.M()
}

package main

import "fmt"

type kloud struct {
	m int
	n string
}

var (
	x = kloud{1, "kloudone"}   // has type kloud
	y = kloud{m: 1}            // Y:0 is implicit
	z = kloud{}                // X:0 and Y:0
	p = kloud{m: 1, n: "tech"} // has type kloud
	// q = kloud{1} causes error too few values in kloud literal
	t = &kloud{1, "devops"} // has type *Vertex

)

func main() {
	fmt.Println(x, t, y, z, p)
}

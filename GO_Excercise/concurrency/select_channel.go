package main

import (
	"fmt"
)

type degen struct {
	i, j string
}

func (x degen) CVIO(ch1, ch2 chan string, quit chan int, m, n string) {
	for {
		select {
		case ch1 <- m:
			fmt.Println(x.i)
		case ch2 <- n:
			fmt.Println("ok")
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	quit := make(chan int)
	x := degen{"goosebump", "ok"}
	go x.CVIO(ch1, ch2, quit, "goosebump", "ok")
}

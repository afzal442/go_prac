package main

import (
	"fmt"
	//"sync"
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
	//var wg sync.WaitGroup
	//wg.Add(1)
	go x.CVIO(ch1, ch2, quit, "goosebump", "ok")
	//wg.Done()

	<-ch1 // Receive from CH1 (allowing "ch1 <- m" in go routine to proceed)
	<-ch2 // Receive from CH2 (allowing "ch2 <- n" in go routine to proceed)

	quit <- 1
	//wg.Wait() // Wait for CVIO to end (which it should do due to above send)
}

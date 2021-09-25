package main

import (
	"fmt"
	"runtime"
)

func check_error(x int) {
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x > 0:
		fmt.Println("x is positive")
	case x == 0:
		fmt.Println("x is zero")
	default:
		fmt.Println("x is unknown")
	}
}

func main() {
	check_error(-5) // x is negative

	fmt.Print("GO is working on ")
	switch os := runtime.GOOS; os {
	case "ubuntu":
		fmt.Println("Ubuntu")
	case "linux":
		fmt.Println("Linux.")
	case "darwin":
		fmt.Println("OS X.")
	case "freebsd":
		fmt.Println("FreeBSD.")
	default:
		fmt.Printf("%s.\n", os) //os is Linux.
	}
}

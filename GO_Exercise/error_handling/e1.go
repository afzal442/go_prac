package main

import (
	"fmt"
)

type MyError struct {
	s string
}

func (e *MyError) Error() string {
	return e.s
}

func Newest(text string) error {
	// c := &MyError{text}
	return &MyError{text}
}

func main() {
	fmt.Println("Hello, playground")
	err := &MyError{"This is an error"}
	x := Newest("Error caught")
	if x != nil {
		fmt.Println(x)
	}
	// err is of type error
	if err != nil {
		fmt.Println(err)
	}
}

// this program shows how to import package quote by which we can access the quote.Go() function
package main

import (
	"fmt"

	"github.com/snapcore/snapd/strutil/shlex"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Glass())
	fmt.Println(shlex.Split(`Hello, Gophers! This symbol will be unescaped`))

}

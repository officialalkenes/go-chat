package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	i := 0
	for i < 10 {
		fmt.Println(quote.Go())
		fmt.Println("Hello World")
		i++
	}
}

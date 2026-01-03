package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Njox")

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	//  ch <- 4
	fmt.Println("sent three values")
}

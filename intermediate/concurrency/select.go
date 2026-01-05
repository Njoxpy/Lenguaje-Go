package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello select concurrency \n")

	test1 := make(chan int, 4)

	test1 <- 1
	test1 <- 2
	test1 <- 2
	test1 <- 3

	select {
	case msg := <-test1:
		fmt.Println("Receibeid from channel 1", msg)
	case msg := <-test1:
		fmt.Println("Received from channel 2: ", msg)
	default:
		fmt.Println("This is the default one :) ")
	}

	fmt.Println(test1)

}

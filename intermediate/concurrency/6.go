package main

import "fmt"

func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 11)
	fmt.Println("testing 6!")

	fmt.Println(ch1, ch2)
}

// CREATE -> SEND -> RECEIVE

package main

import (
	"fmt"
)

func main(){
	ch := make(chan int)

	ch <- 10

	x := <-ch

	fmt.Println(x)
	fmt.Println("Hello from channels")
}

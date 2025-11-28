package main

import "fmt"

func main() {
	fmt.Println("Pointers learning")

	x := 10
	y := &x

	fmt.Println(y)
}

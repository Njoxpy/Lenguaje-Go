package main

import "fmt"

func add(x, y int) (result int) {
	return x + y
	return
}

func main() {
	fmt.Println("Named returns!")
	fmt.Println(add(12, 13))
}

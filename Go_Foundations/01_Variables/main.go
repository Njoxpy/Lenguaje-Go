package main

import "fmt"

var (
	isRegistered        = false
	name         string = "Godbless Nyagawa"
)

func main() {
	fmt.Printf("Type: %T, Value:%v \n", isRegistered, isRegistered)
	fmt.Printf("Type: %T, Value:%v \n", name, name)
}

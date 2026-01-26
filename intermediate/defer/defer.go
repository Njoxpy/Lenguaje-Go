package main

import (
	"fmt"
)

func main(){
	defer fmt.Println("Defer 1 run")
	fmt.Println("Testing defer")

	defer fmt.Println("Defer 2 run")
	fmt.Println("testing defer 2")
	defer greetUser()
}

func greetUser(){
	fmt.Println("Hello, Njox")
}

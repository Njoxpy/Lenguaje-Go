package main

import (
	"fmt"
)

func main(){
	fmt.Println("first executed")
	// added into stack
	defer greetUser()
	defer fmt.Println("added into stack")

	fmt.Println("Second executed")
}

func greetUser(){
	fmt.Println("greet user")
}

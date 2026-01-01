package main

import (
	"fmt"
)

func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()

		defer fmt.Println("hello from defer")

		fmt.Println("hello this is normal")
}


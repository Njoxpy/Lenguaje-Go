package main

import "fmt"

func main() {
	go greetUser()
}

func greetUser() {
	fmt.Println("Hello Njox")
}

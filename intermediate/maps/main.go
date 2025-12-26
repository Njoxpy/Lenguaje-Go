package main

import "fmt"

func main() {
	fmt.Println("Golang maps")

	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	fmt.Println(ages)

	fmt.Println("Alice age: ", ages["alice"])

}

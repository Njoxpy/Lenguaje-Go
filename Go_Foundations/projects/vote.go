package main

import "fmt"

func main() {
	fmt.Println("Welcome to the voting game.")

	var miaka int
	fmt.Printf("Jaza miaka yako: ")
	fmt.Scan(&miaka)

	fmt.Printf("Miaka yako ni: %v", miaka)
	fmt.Println("\n")

	if miaka <= 17 {
		fmt.Println("Huruhusiwi kupiga kula.")
	} else {
		fmt.Println("Unaruhusiwa kupiga kula.")
	}
}

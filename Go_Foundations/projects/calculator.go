package main

import (
	"fmt"
)

func addition(x, y int) int {
	return x + y
}

func subtraction(x, y int) int {
	return x - y
}

func multiplication(x, y int) int {
	return x * y
}

func division(x, y int) int {
	return x / y
}

func main() {
	fmt.Println("Karibu kwenye kikotozi!")

	var operation int

	switch operation {
	case 1:
		fmt.Println(addition(12, 12))
	}
}

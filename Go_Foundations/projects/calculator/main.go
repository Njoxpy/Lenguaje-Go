// go mod init github.com/Njoxpy/calculator
package main

import (
	"fmt"

	"github.com/Njoxpy/calculator/pkg/arithmetic"
)

func main() {
	fmt.Println("Calculator")
	a := 12
	b := 12
	sum := arithmetic.Add(a, b)

	fmt.Printf("Addition: %d + %d = %d\n", a, b, sum)
}

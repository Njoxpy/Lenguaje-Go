package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func addFloat(a float64, b float64) float64 {
	return a + b
}

func sum[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("Generics")

	result := add(12, 12)
	jibu := addFloat(12, 12)

	fmt.Println(result, jibu)

	fmt.Println(sum(5, 10))
	fmt.Println(sum(5.44444444, 10))
}

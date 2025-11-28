package main

import "fmt"

func main() {
	fmt.Println("Learning arrays")

	var numbers [3]int
	numbers[0] = 12
	numbers[1] = 14

	fmt.Println(numbers)
	fmt.Println(numbers[0])

	arr := [3]int{1, 2, 3}
	fmt.Println(arr) // [1 2 3]

}

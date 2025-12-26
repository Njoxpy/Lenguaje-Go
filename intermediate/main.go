package main

import "fmt"

func add(x, y int) (result int, err error) {

	if err != nil {
		return
	}

	result = x + y
	return
}

func main() {
	fmt.Println(add(42, 13))
}

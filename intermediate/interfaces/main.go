package main

import "fmt"

type Cook struct {
	Name string
	Age  int
}

type FoodCook interface {
	Pika()
}

func (c Cook) Pika() {
	fmt.Println(c.Name, " is cooking.")
}
func main() {
	fmt.Println("Interfaces")
}

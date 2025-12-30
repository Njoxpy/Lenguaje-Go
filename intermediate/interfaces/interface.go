package main

import "fmt"

type FoodCook interface {
	Pika()
}

type Cook struct {
	Name string
}

func (c Cook) Pika() {
	fmt.Println(c.Name, "is cooking")
}

func StartCooking(fc FoodCook) {
	fc.Pika()
}

func main() {
	cook := Cook{Name: "Alex"}
	StartCooking(cook)
}

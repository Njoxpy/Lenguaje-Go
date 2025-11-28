package main

import "fmt"

/*
🧠 Exercise 1: Simple Pointer Update

Create a struct called Student with fields Name and Marks.

Create one variable s1 := Student{"Ali", 70}

Make a pointer to it (s2 := &s1)

Update marks to 90 using the pointer.

Print both s1.Marks and s2.Marks — they should be the same.

👉 Goal: understand how pointers share the same memory.
*/

type Students struct {
	name string
	age  int
}

func main() {
	fmt.Println("Learning structs")

	s1 := Students{"Njox", 70}
}

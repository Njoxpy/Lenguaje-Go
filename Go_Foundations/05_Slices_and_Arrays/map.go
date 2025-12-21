package main

import "fmt"

func main() {
	fmt.Println("Map in Go")

	
		students := make(map[string]int){
			"Godbless Nyagawa" : 23,
			"Njox Nyagawa" : 25
		}
	

	s := map[string]int{
		"Njox":  23,
		"Mdudu": 45,
	}

	s["kaka"] = 45
	delete(s, "kaka")

	fmt.Println(len(s))
}

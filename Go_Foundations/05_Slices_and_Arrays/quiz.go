/*
Quiz 1 — Basic creation & access

Create a file quiz1.go.

Create and initialize a map ages with entries: "Alice": 30, "Bob": 25.

Print the age of "Alice".

Read "Charlie" from the map using comma-ok and print whether the key exists.
*/
package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	println("Alice age", ages["Alice"])
	fmt.Println("Charlie exists", ages["Charlie"])

	exists := ages["Charlie"]

	if exists == 0 {
		fmt.Println("does not exists")
	}

	/*
		Quiz 2 — Add / modify / delete

		Start with an empty map inventory := make(map[string]int).

		Add "apple": 10, "banana": 5.

		Increase apple by 5.

		Remove "banana".

		Print final map and its length.
	*/

	inventory := make(map[string]int)

	inventory["apple"] = 10
	inventory["banana"] = 5

	inventory["apple"] += 5

	delete(inventory, "banana")

	fmt.Println(inventory)

	/*

		Quiz 3 — Frequency counter

		Write a function WordCount(s string) map[string]int that returns counts of words separated by spaces.

		Test with "go go gopher go".

		Print top 2 words by count (just iterate and identify).
	*/

	func WordCount(s string) map[string]int{
		
	}

}

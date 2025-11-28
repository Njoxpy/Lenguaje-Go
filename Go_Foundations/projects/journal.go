package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let us rate your day.")
	var score int
	var answer string

	// morning: meditation, exercise, learning, evaluation,
	// user answers what did he get done for today then the system gives score to the user.
	fmt.Println("You did meditation: ")
	fmt.Scan(&answer)

	if answer == "yes" {
		fmt.Println("That is good you did meditation.")
		score += 1
	} else {
		fmt.Println("Make sure you do meditation, next time points will be deducated")
	}

	fmt.Println("You did exercise: ")
	fmt.Scan(&answer)

	if answer == "yes" {
		fmt.Println("That is good you did exercose.")
		score += 1
	} else {
		fmt.Println("Make sure you do exercise, next time points will be deducated")
	}

	fmt.Println("You learned something today: ")
	fmt.Scan(&answer)

	if answer == "yes" {
		fmt.Println("That is good, never stop learning.")
		score += 1
	} else {
		fmt.Println("Don't procrastinate make sure you learn bro")
	}

	fmt.Println("You did day evaluation: ")
	fmt.Scan(&answer)

	if answer == "yes" {
		fmt.Println("Evaluation measures your progress.")
		score += 1
	} else {
		fmt.Println("You need to evaluate your day to know where you are, right")
	}

	results := score * 25

	fmt.Println("You scored: ", results)

}

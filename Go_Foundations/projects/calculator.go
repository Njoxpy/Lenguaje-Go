package main

import (
	"fmt"
)

func main(){
	fmt.Println("Karibu kwenye kikotozi!")

	var namba1 int
	var namba2 int
	var jibu int

	fmt.Printf("Jaza namba ya kwanza: ")
	fmt.Scan(&namba1)

	fmt.Printf("Jaza namba ya pili: ")
	fmt.Scan(&namba2)

	jibu = namba1 + namba2

	fmt.Printf("Jibu: %v\n", jibu)

	jibuAfterSub := namba1 - namba2
	fmt.Printf("Jibu after sub: %v\n", jibuAfterSub)

	jibuAfterMult := namba1 * namba2
	fmt.Printf("Jibu after multiplication: %v\n", jibuAfterMult)

	jibuAfterDivision := namba1 / namba2
	fmt.Printf("Jibu after division: %v\n", jibuAfterDivision)

}

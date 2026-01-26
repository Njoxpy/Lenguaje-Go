
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Defer bucles de programación")

    for i := 0; i <= 10; i++ {
        //      defer
      defer  fmt.Println(i)
    }

		fmt.Println("Loop finished - now returning from main()")
}

/*
- note if the loop contains a defer statement into the loop then code is executed into the reverse order 
*/

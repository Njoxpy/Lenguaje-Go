package main

import (
	"fmt"
)

func main() {
    for i := 0; i < 5; i++ {
        defer func(n int) {
            fmt.Print(n, " ")
        }(i * 2)
    }
}

// 5, 4, 3, 2, 1

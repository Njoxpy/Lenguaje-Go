package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mark this goroutine as finished

	fmt.Println("Worker", id, "is working")
}

func main() {
	var wg sync.WaitGroup

	// we plan to start 3 goroutines
	wg.Add(3)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	// wait until all goroutines call Done()
	wg.Wait()

	fmt.Println("All workers finished")
}


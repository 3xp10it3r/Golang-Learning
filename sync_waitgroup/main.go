package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started.\n", i)
	fmt.Printf("Worker %d ended.\n", i)
}

func main() {

	var wg sync.WaitGroup

	for i:=1; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Workers task completed.")
}

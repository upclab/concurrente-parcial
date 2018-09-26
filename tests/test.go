package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func runSync() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("============== Iteration %d ==============\n", i)
		for j := 1; j <= 8; j++ {
			fmt.Printf("            (%d) Function %d\n", i, j)
		}
		fmt.Printf("============ End Iteration %d ============\n\n", i)
	}
}

func runAsync() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("============== Iteration %d ==============\n", i)

		wg.Add(8)
		for j := 1; j <= 8; j++ {
			go func(jC int) {
				fmt.Printf("            (%d) Function %d\n", i, jC)
				wg.Done()
			}(j)
		}
		wg.Wait()

		fmt.Printf("============ End Iteration %d ============\n\n", i)
	}
}

func main() {
	runAsync()
}

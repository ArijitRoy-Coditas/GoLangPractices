package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d is finished execution\n", id)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		// wg.Add(1)
		go worker(i, &wg)
	}
	// wg.Wait()
	// fmt.Println("All worker finished execution!")
}

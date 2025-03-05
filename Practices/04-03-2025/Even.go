package main

import (
	"fmt"
	"sync"
)

func findEven(num int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	if num%2 == 0 {
		res <- num
	}
}

func aggregate(res <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for num := range res {
		sum = sum + num
	}
	fmt.Println("The sum of the even number is: ", sum)
}
func main() {
	list := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := make(chan int, len(list))
	var wg sync.WaitGroup
	var wgA sync.WaitGroup

	wg.Add(len(list))
	for _, value := range list {
		go findEven(value, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	wgA.Add(1)
	go aggregate(results, &wgA)
	wgA.Wait()

	fmt.Println("Done")
}

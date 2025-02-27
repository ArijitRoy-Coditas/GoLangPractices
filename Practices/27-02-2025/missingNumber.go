package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1}

	n := len(nums)

	n = (n * (n + 1)) / 2

	// fmt.Print(n)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	fmt.Println(n - sum)
}

package main

import(
	"fmt"
	"sort"
)
func main(){

	numbers := []int{5, 3, 8, 1}
	sort.Slice(numbers, func(i, j int) bool {
    	return numbers[i] < numbers[j]
	})

fmt.Println(numbers)

}
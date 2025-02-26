package main

import (
	"fmt"
)

func main() {
	hashMap := make(map[int]int)

	nums := []int{2,2,1,1,1,2,2}
	maxValues := 0
	maxKey := nums[0]
	for _,values := range nums {
		hashMap[values]++
	}
	fmt.Println(hashMap)

	for key, values := range hashMap{
		if values > maxValues {
			maxValues = values
			maxKey = key
		}
	}

	fmt.Println(maxKey)
}

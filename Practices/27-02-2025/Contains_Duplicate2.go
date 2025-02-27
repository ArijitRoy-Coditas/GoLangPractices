package main

import (
	"fmt"
)

func Check() bool {
	nums := []int{1, 2, 3, 1, 2, 3}
	hashMap := make(map[int]int)
	k := 2

	for key, values := range nums {
		if index, ok := hashMap[values]; ok {
			if key-index <= k {
				return true
			}
		}
		hashMap[values] = key
	}
	fmt.Println(hashMap)
	return false
}

func main() {
	fmt.Println(Check())
}

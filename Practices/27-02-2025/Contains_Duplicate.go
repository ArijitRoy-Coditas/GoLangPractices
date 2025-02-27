package main

import "fmt"

func duplicate() bool {
	nums := []int{1, 4, 3, 2}
	hashMap := make(map[int]int)

	for key, values := range nums {
		if _, ok := hashMap[values]; ok {
			return true
		}
		hashMap[values] = key
	}
	return false
}

func main() {
	fmt.Println(duplicate())
}

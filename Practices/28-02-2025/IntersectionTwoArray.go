package main

import "fmt"

func main() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	hashMap := make(map[int]int)
	for keys, values := range num2 {
		hashMap[values] = keys
	}
	fmt.Println(hashMap)
	res := []int{}
	for _, values := range num1 {
		if _, exists := hashMap[values]; exists {
			res = append(res, values)
			delete(hashMap, values)
		}
	}

	fmt.Println(res)
}

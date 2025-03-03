package main

import "fmt"

func main() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	hashMap := make(map[int]int)
	res := []int{}
	for _, values := range num2 {
		hashMap[values]++
	}

	for _, values := range num1 {
		if _, ok := hashMap[values]; ok && hashMap[values] >= 1 {
			res = append(res, values)
			hashMap[values]--
		}
	}

	fmt.Println(res)
}

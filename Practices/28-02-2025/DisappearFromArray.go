package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	num := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// hashMap := make(map[int]int)

	res := []int{}
	for keys := 0; keys < len(num); keys++ {
		index := abs(num[keys]) - 1
		if num[index] > 0 {
			num[index] = -num[index]
		}
	}
	fmt.Println(num)
	for keys := 0; keys < len(num); keys++ {
		if num[keys] >= 0 {
			res = append(res, keys+1)
		}
	}
	fmt.Println(res)

	// for i := 0; i < len(num); i++ {
	// 	if _, ok := hashMap[num[i]]; ok {
	// 		hashMap[num[i]]--
	// 		fmt.Println(i)
	// 	}
	// }
	// fmt.Println(hashMap)

}

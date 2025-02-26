package main

import "fmt"

func main() {

	nums := []int{3, 1, 3}
	hashMap := make(map[int]int)
	count := 0
	for _, values := range nums {
		count++
		if _, exists := hashMap[values]; !exists {
			fmt.Println("Inside",count)
			hashMap[values] = count
			fmt.Println(hashMap[values])
		}
		// hashMap[values] = count
		// fmt.Println(key, values)
	}

	// for key, values := range hashMap {
	// 	fmt.Println(values, hashMap[key])
	// }

}
package main

import "fmt"



// type HashMap struct {
// 	table map[int]int
// }

// func NewHashTable() *HashMap {
// 	return &HashMap{
// 		table: make(map[int]int),
// 	}
// }


// func (h *HashMap) Add(key, value int) {
// 	h.table[key] = value
// }

// func (h *HashMap) Get(key int) int {
// 	return h.table[key]
// }

func main() {
	array := []int{2,3,8,7}

	target := 9

	hashMap := make(map[int]int)

	for key, nums := range array {

		if index, exists := hashMap[target - nums]; exists {
			fmt.Println("The indexs are: ",index,key)
		}

		hashMap[nums] = key
	}


}
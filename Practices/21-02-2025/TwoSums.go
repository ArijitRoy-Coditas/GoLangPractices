package main

import "fmt"

type HashMap struct {
	table map[int]int
}

func NewHashTable() *HashMap {
	return &HashMap{
		table: make(map[int]int),
	}
}

func (h *HashMap) Add(key, value int) {
	h.table[key] = value
}

func (h *HashMap) Get(key int) int {
	return h.table[key]
}

func main() {
	nums := []int{2, 7, 8, 5}
	target := 9
	h := NewHashTable()
	for key, num := range nums {
		if index, exists := h.table[target-num]; exists {
			fmt.Printf("%d,%d\n",index,key)
		}
		h.Add(num, key)

	}
}
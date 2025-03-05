package main

import "fmt"

func nearestGreaterElement(num1, num2 []int) []int {
	nextGreaterMap := make(map[int]int)
	stack := []int{}

	for _, num := range num2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreaterMap[top] = num
		}
		stack = append(stack, num)
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nextGreaterMap[top] = -1
	}

	result := make([]int, len(num1))
	for i, num := range num1 {
		result[i] = nextGreaterMap[num]
	}
	return result
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nearestGreaterElement(nums1, nums2))
}

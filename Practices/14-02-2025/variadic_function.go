package main 

import "fmt"

func sum(nums ...int) int{
	var num int
	for i := range nums {
		num += nums[i]
	}
	return num
}

func main() {
	a := sum(1,2,3)
	fmt.Println(a)
}
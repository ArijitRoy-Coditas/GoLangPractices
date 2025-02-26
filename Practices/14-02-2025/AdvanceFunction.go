package main

import "fmt"

func sum(a, b int) int{
	return a + b
}

func mul(a, b int) int {
	return a * b
}
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a , b), c)
}
func main(){
	fmt.Println(aggregate(2,3,4, sum))

	fmt.Println(aggregate(2,4,6, mul))
}
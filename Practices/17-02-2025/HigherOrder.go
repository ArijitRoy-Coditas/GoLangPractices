package main

import (
	"fmt"
)


func multiply(a, b int) int {
	return a * b
}


func selfMath(Math func(a, b int) int) func(int) int{
	return func(x int) int {
		return Math(x,x)
	}
}

func main(){
	SquareFunc := selfMath(multiply)
	fmt.Println(SquareFunc(5))
}

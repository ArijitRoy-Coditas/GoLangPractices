package main 

import "fmt"

func main() {
	// var myInt = [...]int {1,2,2,4,5}

	// myInt = {1,2,2,4,5,}
	myInt := []int{1,2,3,4,5}

	b := myInt[1:4]

	for i := range b {
		fmt.Println(b[i])
	}

	fmt.Println(b[:]) //Prints the all value
	fmt.Println(b[1:]) //Prints from the index 1
	fmt.Println(b[:1]) //Prints the first value only since the right side of colon is exclusive

	// c := "Hello"


	/* Slicing is not same as python it is much simpler with start and end value. 
	start value is inclusive and end value is not inclusive */

	// fmt.Println(c[::-1]) 
}
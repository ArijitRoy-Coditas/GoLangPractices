package main

import(
	"fmt"
)

func palindrom(n int) int {
	if(n == 0 || n == 1){
		return n
	}


	// Anothere Correct way
	// if n:= 2; n<1{

	// }

	return palindrom(n - 1) + palindrom(n - 2)
}

func main(){

	//correct way
	// for i := 0; i<8; i++{
	// 	fmt.Println(palindrom(i))
	// }

	//correct way
	for i:=range 8{
		fmt.Println(palindrom(i))
	}

	//short declaration of the variable.
	a := 12
	b := 15

	//swap works same as python
	a, b = b,a
	fmt.Println(a)
	fmt.Println(b)
}
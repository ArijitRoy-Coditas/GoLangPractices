package main

import(
	"fmt"
)

func check(sum int) (x int, y float32){
	x = sum * 4 / 9
	y = float32(sum / 4)

	return x, y
}

func main(){
	fmt.Println(check(9))
}
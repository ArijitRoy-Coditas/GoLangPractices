package main

import "fmt"

func pal(x int) bool {
	sum := 0
	newNum := x

	for x > 0 {
		sum = sum * 10 + (x % 10)
		x = x/10

	}

	return newNum == sum
}

func main() {
	if pal(121) {
		fmt.Println("Pal")
	} else {
		fmt.Println("Not Pal")
	}
}
package main

import (
	"fmt"
	"strconv"
)

func reverse[T ~int](a T) bool {
	revStr := []rune(strconv.Itoa(int(a)))

	for i, j := 0, len(revStr) - 1; i<j; i,j = i + 1, j - 1 {
		revStr[i], revStr[j] = revStr[j], revStr[i]
	}

	if strconv.Itoa(int(a)) == string(revStr) {
		return true
	}
	return false 
}

func main() {
	num := 0
	fmt.Println("Enter the number you want to check for palindrome: ")
	fmt.Scanln(&num)
	if reverse(num) {
		fmt.Println("The given number is a palindrome")
	} else {
		fmt.Println("The given number is not a palindrome")
	}
}
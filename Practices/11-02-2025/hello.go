package main

import "fmt"

func main() {
	fmt.Println("Hello")
	const str1, str2, str3 = "a","b","c"
	s := fmt.Sprint(str1,str2,str3)
	fmt.Print(s)
	var str4, str5 = "Arijit", "Roy"
	x := fmt.Sprintln(str4,str5)
	fmt.Println()
	fmt.Println(x)
}

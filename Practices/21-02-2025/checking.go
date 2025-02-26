package main

import (
	"fmt"
)

func main() {
	s := "Arijit"
	
	// a := make([]string, 10)
	for _, a := range s {
		fmt.Printf("%c\n",a)
	}
}
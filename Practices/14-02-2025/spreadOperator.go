package main 

import "fmt"

func printString(names ...string){

	for i:= range names {
		fmt.Println(names[i])
	}
}

func main(){
	names := []string{"bob", "sue", "alice"}
	printString(names...)
}
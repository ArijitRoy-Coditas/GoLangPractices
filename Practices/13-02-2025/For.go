package main


import "fmt"


var tring string = "Hello how are you doing"

func main(){
	for i := range tring {
		// fmt.Println(i)
		fmt.Print(string(tring[i]))
	}
}
package main

import(
	"fmt"
	"math/rand"
	"math"
)
func main(){
	fmt.Println(rand.Intn(1000)) //Gives a random number between 0 - 1000
	fmt.Println(math.MaxInt16) // Shows the maximum possbile value that can be assigned in 16-bit
	fmt.Println(math.MaxInt8) // Similar to 16-bit shows the maximum possible value for 8-bit
	// var a = 1231312424241123454 // Just to check how large value can be stored.
	var a int8 = 127 // Maximum value possible in int8
	var c int8 = -128 //Minimum value possible in int8
	fmt.Println(a)
	fmt.Println(c)
	var d int64 = 9223372036854775808
	fmt.Println(d)
	fmt.Println(math.MaxInt64)

	fmt.Printf("The square %d root is: %g", 25, math.Sqrt(25)) //%g removes any decimal place or extra zero
	fmt.Println()
	fmt.Printf("%.3g",math.Pi)
}
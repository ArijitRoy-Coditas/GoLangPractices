package main

import "fmt"

func main() {
fmt.Println("---Example 1---")
// Example 1
// define s source slice
  sourceSlice := []int{11,22,33}
// define a destination slic
  destSlice := []int{44,55,66,77}
// call copy function 
  numberCopied := copy(destSlice, sourceSlice)
// display the destination slice
  fmt.Println(destSlice)
// display the number of copied elements
  fmt.Println("Number of elements copied:")

  fmt.Println(numberCopied)

  fmt.Println("---Example 2---")

// Example 2

// use the same slice as the source and destination slice
  numberCopied = copy(destSlice, destSlice[2:])
// display the destination slice
  fmt.Println(destSlice)
// display the number of copied elements
  fmt.Println("Number of elements copied:")
  fmt.Println(numberCopied)
}
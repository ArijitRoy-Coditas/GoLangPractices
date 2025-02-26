package main

import(
	"fmt"
)

func add() func (a int) int{
	sum := 0
	return func (a int) int {
		sum += a
		return sum
	}
}

func main() {
	Addition := add()
	Addition(5)
	Addition(5)
	Addition(5)
	Addition(5)
	Addition(5)
	Addition(5)
	fmt.Println(Addition(5))
}
package main

import "fmt"

func printReports(intro, body, outro string) {
	printCostReport(func(a string) int {
		return len(a) * 2
	}, intro)

	printCostReport(func (a string) int{
		return len(a) * 3
	}, body)

	printCostReport(func (a string) int{
		return len(a) * 4
	}, outro)
	
}

// don't touch below this line

func main() {
	printReports(
		"Such a lovely place",
		"Welcome to the Hotel California",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
package main

import (
	"fmt"
)

func concatter() func (string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}
func main() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Hi")
	harryPotterAggregator("How are you?")
	harryPotterAggregator("What is your name")
	fmt.Println(" ",harryPotterAggregator("My name is Arijit Roy"))
}
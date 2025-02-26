package main

import (
	"fmt"
)

type messageToSend struct{
	phoneNumber int
	message string
} 


func test (m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}


func main() {
	test(messageToSend{
		phoneNumber: 8911231231,
		message: "Thanks for signing up",
	})
}
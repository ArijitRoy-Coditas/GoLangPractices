package main 

import (
	"fmt"
	"time"
)

func sendEmail(message string) {

	func() {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("Email recieved %s\n",message)
	}()
	
	fmt.Printf("Email send %s\n",message)
}

func test(message string){
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("=====================================")
}

func main() {

	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}
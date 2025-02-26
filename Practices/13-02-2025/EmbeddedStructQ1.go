package main

import (
	"fmt"
)

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	//Create the instance like this to access the User or to make change in the variables inside the Membership

	// user := User{Name: name}

	//Alternatively you can use the below method by creating a variable like u User.
	var u User

	//Both the above method correct.
	u.Name = name
	if membershipType == "Premium" {
		u.Type = "Premium"
		u.MessageCharLimit = 1000
	} else {
		u.Type = "Standard"
		u.MessageCharLimit = 100
	}

	return u
}


//Another alternative way to do this

// func newUser(name string, membershipType string) User {
// 	membership := Membership{Type: membershipType}
// 	if membershipType == "premium" {
// 		membership.MessageCharLimit = 1000
// 	} else {
// 		membership.Type = "standard"
// 		membership.MessageCharLimit = 100
// 	}
// 	return User{Name: name, Membership: membership}
// }

func (u User) sendMessage(message string, messageLength int,) (string, bool) {

	if messageLength <= u.MessageCharLimit {
		return message, true
	}

	return "", false
}

func main(){
	user1 := newUser("Arijit", "Premium")

	fmt.Printf("Name: %s\nMessageCharLimit: %d\nMembershipType: %s", user1.Name, user1.MessageCharLimit, user1.Type)
	fmt.Println()

	message1, canSend1 := user1.sendMessage("Hello, this is a test message.", len("Hello, this is a test message."))
	if canSend1 {
		fmt.Printf("Message Sent: %s\n", message1)
	} else {
		fmt.Println("Message exceeds character limit.")
	}

	fmt.Println()
	fmt.Printf("Name: %s\nMessageCharLimit: %d\nMembershipType: %s", newUser("Jit", "Standard").Name, newUser("Jit", "Standard").MessageCharLimit, newUser("Jit", "Standard").Type)
}
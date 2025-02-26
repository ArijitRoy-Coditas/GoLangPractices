package utils

import (
	"fmt"
)


func WaitMenu() {

	var option int
	fmt.Println("Please 5 to go back to the menu page")
	fmt.Scanln(&option)

	if option == 5 {
		return
	} else {
		fmt.Println("Invalid input. Please try again")
		WaitMenu()
	}

}
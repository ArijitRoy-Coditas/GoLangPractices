package utils

import (
	"Todo_list/ModifyOption"
	"fmt"
	"time"
)

func DisplayModifyMenu(returnToMenuCallback func(), Id int) {

	fmt.Println("Option 1: Modify the Task")
	fmt.Println("Option 2: Modify the Description")
	fmt.Println("Option 3: Modify the Status")
	fmt.Println("Option 4: Press 5 to go back to the Main menu")
	fmt.Println("=================================================")
	fmt.Println()

	var options int

	fmt.Printf("Enter your choice: \n")
	fmt.Scanln(&options)

	if options == 5 {
		returnToMenuCallback()
	}

	switch options {
	case 1:
		ClearConsole()
		modifyoption.ModifyTask(Id)
		LoadingPrompt("Hang on!, We are updating the task...", time.Millisecond * 1000)
		fmt.Println("Task has been updated successfully")
		fmt.Println("==================================================")
	case 2:
		ClearConsole()
		modifyoption.ModifyDescription(Id)
		LoadingPrompt("Hang on!, We are updating the description...", time.Millisecond * 1000)
		fmt.Println("Description has been updated successfully")
		fmt.Println("==================================================")
	case 3:
		ClearConsole()
		modifyoption.ModifyStatus(Id)
		LoadingPrompt("Hang on!, We are updating the status...", time.Millisecond * 1000)
		fmt.Println("Status has been updated successfully")
		fmt.Println("==================================================")
	}
}

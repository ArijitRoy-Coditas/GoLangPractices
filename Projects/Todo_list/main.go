package main 

import (
	"Todo_list/utils"
	"Todo_list/Actions"
	"fmt"
	"os"
	"time"
)


func DisplayMenu() {

	for {

		utils.ClearConsole()

		fmt.Println("Todo List Menu")
		fmt.Println("Option 1: Read Todo List")
		fmt.Println("Option 2: Add New Todo")
		fmt.Println("Option 3: Modify Todo")
		fmt.Println("Option 4: Delete Todo List")
		fmt.Println("Press any key to exit...")
		fmt.Println("=================================================")
		fmt.Println()

		var options int //Logical name: choose variable carefully

		fmt.Printf("Enter your choice: \n")
		fmt.Scanln(&options)

		switch options {
		case 1:
			utils.ClearConsole()
			utils.LoadingPrompt("Hang on! we are fetching your details.....", time.Millisecond * 1000)
			Actions.ReadList("Successfully Read\n=============================================")
			utils.WaitMenu()

		case 2:
			utils.ClearConsole()
			Actions.AddTask()
			utils.WaitMenu()
		case 3:
			utils.ClearConsole()
			Actions.Modify(func() {
				DisplayMenu()
			})
			utils.WaitMenu()
		case 4:
			Actions.Delete()
		default:
			fmt.Println("Thank you for visiting the Todo List....")
			os.Exit(0)
		}
	}

}

func main() {
	DisplayMenu()
}
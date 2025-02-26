package Actions

import (
	"Todo_list/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"github.com/olekukonko/tablewriter"
	"time"
)

func Modify(returnToMenuCallback func()) {

	utils.LoadingPrompt("Hang on! we are fetching your details.....", time.Millisecond * 1000)
	ReadList("Please select the ID for the task you want to modify...")

	var Id int
	fmt.Scanln(&Id)

	flag := os.O_APPEND | os.O_RDWR
	file, err := os.OpenFile("./data/tasks.csv", flag, 0644)
	if err != nil {
		fmt.Println("Error opening the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading the file", err)
	}

	utils.ClearConsole()
	utils.LoadingPrompt("Hang on! We are generating options for you.....", time.Millisecond * 1500)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Task", "Description", "Status", "CreatedAt"})
	for _, eachRecords := range records {
		if eachRecords[0] == strconv.Itoa(Id) {

			utils.ClearConsole()

			table.Append(eachRecords)

		}
	}

	table.Render()
	utils.DisplayModifyMenu(returnToMenuCallback, Id)

}

package Actions

import (
	// "Todo_list/utils"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	// "strconv"
	"github.com/olekukonko/tablewriter"

	
)


func ReadList(message string){

	file, err := os.Open("./data/tasks.csv")
	if err != nil {
		log.Println("Error opening the file.",err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Error reading the file.",err)
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"ID", "Task", "Description", "Status", "CreatedAt"})

	for _, eachRecords := range records[1:] {
		table.Append(eachRecords)
	}
	table.Render()

	fmt.Println("=============================================")
	fmt.Println(message)

}
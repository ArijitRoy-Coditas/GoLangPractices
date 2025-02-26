package modifyoption

import (
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"strconv"
)

func ModifyStatus(Id int) {
	flag := os.O_APPEND | os.O_RDWR
	file, err := os.OpenFile("./data/tasks.csv", flag, 0644)
	if err != nil {
		fmt.Println("Error opening the file",err)
	}

	reader := csv.NewReader(file)

	recorder, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading the file", err)
	}

	var newTask string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the new status that you want to add: ")
	scanner.Scan()

	newTask = scanner.Text()

	for key, eachRecord := range recorder {
		
		if eachRecord[0] == strconv.Itoa(Id) {
			eachRecord[3] = newTask
			recorder[key] = eachRecord
			break
		}
	}

	if err := file.Close(); err != nil {
		fmt.Println("Error closing the file:", err)
		return
	}	

	file, err = os.OpenFile("./data/tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644 )
	if err != nil {
		fmt.Println("Error opening the file", err)
	}

	writer := csv.NewWriter(file)

	for _, eachRecords := range recorder {
		writer.Write(eachRecords)
	}

	writer.Flush()
	defer file.Close()
}
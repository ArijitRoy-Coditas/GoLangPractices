package modifyoption

import (
	"fmt"
	"os"
	"bufio"
	"encoding/csv"
	"strconv"
)

func ModifyDescription(Id int) {

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
	fmt.Println("Enter the new description that you want to add: ")
	scanner.Scan()

	newTask = scanner.Text()

	for i, eachRecord := range recorder {
		
		if eachRecord[0] == strconv.Itoa(Id) {
			eachRecord[2] = newTask
			recorder[i] = eachRecord
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
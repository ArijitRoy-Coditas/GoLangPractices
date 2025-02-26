package Actions

import (
	"Todo_list/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Delete() {
	
	utils.LoadingPrompt("Hang on! we are fetching your details.....", time.Millisecond * 1000)
	ReadList("Please select the ID for the task you want to delete...")

	var Id int
	fmt.Scanln(&Id)

	flag := os.O_RDONLY
	file, err := os.OpenFile("./data/tasks.csv", flag, 0644)
	if err != nil {
		fmt.Println("Error opening the file",err)
	}

	reader := csv.NewReader(file)
	recorder, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading the file",err)
	}

	newRecorder := [][]string{}

	for _, eachRecord := range recorder {
		
		if eachRecord[0] != strconv.Itoa(Id) {
			newRecorder = append(newRecorder, eachRecord)
		}

	}

	file.Close()

	file, err = os.OpenFile("./data/tasks.csv", os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening the file", err)
	}

	writer := csv.NewWriter(file)
	writer.WriteAll(newRecorder)
	writer.Flush()
	defer file.Close()

	fmt.Println("Successfully read")
}
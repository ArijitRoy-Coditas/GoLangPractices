package Actions

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"Todo_list/utils"
)

type task struct {
	id string
	Task string
	description string
	status string
	createdAt string
}

// Get the highestID present in the file
func GetHighestId(file *os.File) int {
	
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if len(records) <= 1 {
		return 0
	}

	if  lastRecord := records[len(records) - 1]; lastRecord[len(lastRecord) - 1] != "\n" {
		writer := csv.NewWriter(file)
		writer.Write([]string{})
	}

	if err != nil {
		fmt.Println("Error reading the file", err)
	}

	lastRecord := records[len(records) - 1]
	highestId, err := strconv.Atoi(lastRecord[0])
		
	if err != nil {
		fmt.Println(err)
	}

	return highestId
}

// Get the current time
func GetCurrentTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05") //"2006-01-02 15:04:05" default format go uses for converting the time to human readable time
	return formattedTime
}

func (t task) ToSlice() []string {
	return []string{t.id, t.Task, t.description, t.status, t.createdAt}
}

//Write in the csv file
func WriteCSV(T task, file *os.File) {

	writer := csv.NewWriter(file)
	writer.Write(T.ToSlice())
	writer.Flush()
}


// Add new task to the file
func AddTask(){

	// Array or Slice Literal
	a := []string{"Task", "Description"}
	var inputStorage = []string{}

	scanner := bufio.NewScanner(os.Stdin)

	flag := os.O_APPEND | os.O_RDWR

	file, err := os.OpenFile("./data/tasks.csv", flag, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var Task task

	for _, value := range a {
		fmt.Printf("Enter the %s you want to add: ",value)
	
		scanner.Scan()
		input := scanner.Text()

		inputStorage = append(inputStorage, input)		
		fmt.Println("You have entered", input)
	}

	Task.id = strconv.Itoa(GetHighestId(file) + 1)
	Task.Task = inputStorage[0]
	Task.description = inputStorage[1]
	Task.status = "Incomplete"
	Task.createdAt = GetCurrentTime()

	WriteCSV(Task,file)
	utils.LoadingPrompt("=============================================\nAdding your details, please wait.....", time.Millisecond * 1500)
	fmt.Println("Task added succesfully")
	fmt.Println("=============================================")
}
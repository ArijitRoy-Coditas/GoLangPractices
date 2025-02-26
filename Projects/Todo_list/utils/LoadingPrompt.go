package utils

import (
	"fmt"
	"time"
)

func LoadingPrompt(message string, delay time.Duration) {

	fmt.Println(message)
	time.Sleep(delay)
	ClearConsole()

}
package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskManager struct {
	wg    sync.WaitGroup
	tasks sync.Map
}

var taskManager TaskManager

func handlePostTask(ctx *gin.Context) {
	var Task struct {
		TaskId   uint `json:"task_id"`
		Duration uint `json:"duration"`
	}

	if err := ctx.BindJSON(&Task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Response"})
	}

	taskManager.tasks.Store(Task.TaskId, "running")
	taskManager.wg.Add(1)

	go func(taskId uint, duration uint) {
		defer taskManager.wg.Done()
		time.Sleep(time.Duration(Task.Duration) * time.Second)
		taskManager.tasks.Store(Task.TaskId, "Completed")
	}(Task.TaskId, Task.Duration)

	ctx.IndentedJSON(http.StatusAccepted, gin.H{"message": "Task has started", "Task ID": Task.TaskId})
}

func handleGetTaskByID(ctx *gin.Context) {
	idStr := ctx.Param("id")                    // Gin provides task ID as a string
	id, err := strconv.ParseUint(idStr, 10, 64) // Convert to uint
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID"})
		return
	}

	if status, exists := taskManager.tasks.Load(uint(id)); exists {
		ctx.IndentedJSON(http.StatusOK, gin.H{"taskID": id, "status": status})
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
}

func main() {
	server := gin.Default()

	server.POST("/post", handlePostTask)
	server.GET("/post/:id", handleGetTaskByID)

	server.Run(":8080")
}

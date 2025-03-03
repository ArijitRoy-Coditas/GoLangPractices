package handlers

import (
	"Task_manager/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCurrentTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}
func CreateTask(ctx *gin.Context) {
	var newTask models.Tasks

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
		return
	}

	if newTask.Priority < 1 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Priority cannot be negative or zero"})
		return
	}

	newTask.ID = uuid.New()
	newTask.CreatedAT = time.Now()
	newTask.UpdatedAT = time.Now()
	models.Task = append(models.Task, newTask)
	ctx.IndentedJSON(http.StatusCreated, newTask)

}

package handlers

import (
	"Task_manager/models"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	parseId, err := uuid.Parse(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid ID format"})
		return
	}
	var newTask models.Tasks

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "something went wrong"})
	}

	for i := range models.Task {
		if parseId == models.Task[i].ID {
			if newTask.Name != "" {
				models.Task[i].Name = newTask.Name
			}
			if newTask.Description != "" {
				models.Task[i].Description = newTask.Description
			}
			if newTask.Status != "" {
				models.Task[i].Status = newTask.Status
			}
			if newTask.Priority != 0 {
				models.Task[i].Priority = newTask.Priority
			}
			models.Task[i].UpdatedAT = time.Now()
			ctx.IndentedJSON(http.StatusAccepted, gin.H{
				"message": "Task updated successfully",
				"task":    models.Task[i],
			})
			return

		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
}

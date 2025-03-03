package handlers

import (
	"Task_manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTask(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, models.Task)
}

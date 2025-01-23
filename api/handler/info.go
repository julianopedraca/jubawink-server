package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Server Info
// @Description Check if server is returning Ok.
// @Tags Health Check
// @Accept json
// @Success 200 {object} SignupResponse
// @Failure 400 {object} ErrorResponse
// @Router /info [get]
func Info(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is up."})
}

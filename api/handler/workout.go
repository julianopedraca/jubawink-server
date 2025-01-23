package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/internal/database/models"

	log "github.com/sirupsen/logrus"
)

type WorkoutResponse struct {
	Message string `json:"message"`
}

// @BasePath /api/v1

// @Summary Get Workout By UserId
// @Description Fetches workouts for a specific user from the database.
// @Tags Workout
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} WorkoutResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout/user [get]
func GetWorkoutsByUserId(context *gin.Context) {
	userId, exists := context.Keys["userId"].(int64)

	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in request context."})
		return
	}

	var workout models.Workout
	workouts, err := workout.GetWorkoutsByUserId(userId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to fetch workouts from the database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"workouts": workouts})
}

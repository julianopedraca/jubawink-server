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

	var workout models.WorkoutGetByUserId
	workout.UserId = userId
	workouts, err := workout.GetWorkoutsByUserId()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to fetch workouts from the database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"workouts": workouts})
}

// @Summary Save Workout
// @Description Saves a workout for a specific user into the database.
// @Tags Workout
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param workout body models.WorkoutSave true "Workout save details"
// @Success 200 {object} WorkoutResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout/save [post]
func SaveWorkout(context *gin.Context) {
	userId, exists := context.Keys["userId"].(int64)
	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in request context."})
		return
	}

	var workout models.WorkoutSave
	err := context.ShouldBindJSON(&workout)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to bind exercise json.")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	err = workout.Save(userId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to fetch workouts from the database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"workouts": "Workout saved"})
}

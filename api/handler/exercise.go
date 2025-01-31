package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/internal/database/models"

	log "github.com/sirupsen/logrus"
)

type ExerciseResponse struct {
	Message string `json:"message"`
}

// @BasePath /api/v1

// @Summary Add Exercise
// @Description Adds a new exercise to the database.
// @Tags Exercises
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param exercise body models.ExerciseLiftingSave true "Exercise Lifting Save details"
// @Success 200 {object} ExerciseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /exercise/add/lifting [post]
func AddExerciseLifting(context *gin.Context) {
	var exercise models.ExerciseLiftingSave

	err := context.ShouldBindJSON(&exercise)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to bind exercise json.")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	err = exercise.SaveLifting()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to save exercise to database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Exercise saved"})
}

// @BasePath /api/v1

// @Summary Get Exercise
// @Description Get lifting exercises from userId.
// @Tags Exercises
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ExerciseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /exercise/get/lifting [get]
func GetExerciseLifting(context *gin.Context) {
	userId, exists := context.Keys["userId"].(int64)

	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in request context."})
		return
	}

	var exerciseLifting models.ExerciseLiftingGet
	exerciseLifting.UserId = userId
	exercisesLifting, err := exerciseLifting.GetLifting()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to fetch workouts from the database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"exerciseLifting": exercisesLifting})
}

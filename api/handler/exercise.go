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

// @Summary Add Cycling Exercise
// @Description Adds a new cycling exercise to the database.
// @Tags Exercises
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param exercise body models.ExerciseCyclingSave true "Exercise cycling Save details"
// @Success 200 {object} ExerciseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /exercise/add/cycling [post]
func AddExerciseCycling(context *gin.Context) {
	var exercise models.ExerciseCyclingSave

	err := context.ShouldBindJSON(&exercise)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to bind exercise json.")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	err = exercise.SaveCycling()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to save exercise to database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Exercise saved"})
}

// @Summary Get Exercise
// @Description Get cycling exercises from userId.
// @Tags Exercises
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ExerciseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /exercise/get/cycling [get]
func GetExerciseCycling(context *gin.Context) {
	userId, exists := context.Keys["userId"].(int64)

	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in request context."})
		return
	}

	var exerciseCycling models.ExerciseCyclingGet
	exerciseCycling.UserId = userId
	exercisesLifting, err := exerciseCycling.GetCycling()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to fetch workouts from the database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"exerciseLifting": exercisesLifting})
}

// @Summary Add Running Exercise
// @Description Adds a new Running exercise to the database.
// @Tags Exercises
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param exercise body models.ExerciseRunningSave true "Exercise Running Save details"
// @Success 200 {object} ExerciseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /exercise/add/running [post]
func AddExerciseRunning(context *gin.Context) {
	var exercise models.ExerciseRunningSave

	err := context.ShouldBindJSON(&exercise)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to bind exercise json.")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data. Please check the fields and try again."})
		return
	}

	err = exercise.SaveRunning()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to save exercise to database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save exercise. Please try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Exercise saved successfully."})
}

// @Summary Get Running Exercises
// @Description Retrieves all running exercises for a specific user.
// @Tags Exercises
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ExerciseResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /exercise/get/running [get]
func GetExerciseRunning(context *gin.Context) {
	userId, exists := context.Keys["userId"].(int64)

	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in request context."})
		return
	}

	var exerciseCycling models.ExerciseRunningGet
	exerciseCycling.UserId = userId
	exercisesLifting, err := exerciseCycling.GetRunning()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to fetch workouts from the database.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"exerciseLifting": exercisesLifting})
}

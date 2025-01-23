package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/internal/database/models"
	"github.com/julianopedraca/jubawink/pkg/utils"

	log "github.com/sirupsen/logrus"
)

type SignupResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// @BasePath /api/v1

// @Summary User Signup
// @Description Creates a new user account with the provided details. Passwords are hashed before being stored.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 200 {object} SignupResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /signup [post]
func Signup(ginContext *gin.Context) {
	var user models.User
	err := ginContext.ShouldBindJSON(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to bind user json.")
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format."})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable hash password.")
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	user.Password = hashedPassword

	err = user.Save()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to save user to the data base.")
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to send email.")
		ginContext.JSON(http.StatusBadGateway, gin.H{"error": "Unable to send email, please try again later."})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "User created."})
}

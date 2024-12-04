package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/internal/database/models"
	"github.com/julianopedraca/jubawink/pkg/utils"

	log "github.com/sirupsen/logrus"
)

type LoginResponse struct {
	Token string `json:"token"`
}

// @BasePath /api/v1

// @Summary User Login
// @Description Authenticates a user by validating email and password, and returns a JWT token upon successful login.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userCredentials body models.UserCredentials true "User email and password"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func Login(context *gin.Context) {
	var user models.UserCredentials
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to bind user json.")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	userCredentials, userId, err := user.FindUserByEmail()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to find email user on database.")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email or password invalid."})
		return
	}

	isValidLogin := utils.CheckPasswordHash(userCredentials.Password, user.Password)
	if !isValidLogin {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email or password invalid."})
		return
	}

	token, err := utils.GenerateJwtToken(userId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to generate jwt token.")
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": token})
}

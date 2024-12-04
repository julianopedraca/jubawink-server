package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/pkg/utils"

	log "github.com/sirupsen/logrus"
)

type AuthResponse struct {
	Message string `json:"message"`
}

type Auth struct {
	Token string `json:"token" binding:"required"`
}

// @BasePath /api/v1

// @Summary Validate Token
// @Description Validates a given JWT token to ensure it is authentic and not expired.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param auth body Auth true "Token to validate"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/validate [post]
func ValidateToken(context *gin.Context) {
	var auth Auth

	err := context.ShouldBindJSON(&auth)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to bind auth json.")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	_, err = utils.VerifyToken(auth.Token)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to verify token.")
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid token."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "valid token."})
}

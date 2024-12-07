package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/internal/database/models"
	"github.com/julianopedraca/jubawink/internal/redis"
	"github.com/julianopedraca/jubawink/pkg/mail"
	"github.com/julianopedraca/jubawink/pkg/utils"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

type SignupResponse struct {
	Message string `json:"message"`
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
	var rdb = redis.Rdb
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

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

	uuid := uuid.NewString()
	user.Password = hashedPassword

	userData := &models.User{
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}

	userJson, err := json.Marshal(userData)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to marshal user data.")
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	err = rdb.JSONSet(ctx, uuid, "$", userJson).Err()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to save user to redis.")
		ginContext.JSON(http.StatusBadGateway, gin.H{"error": "Unable to save user to redis."})
		return
	}

	err = mail.SendSignupConfirmation(user.Email, uuid)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to send email.")
		ginContext.JSON(http.StatusBadGateway, gin.H{"error": "Unable to send email, please try again later."})
		return
	}

	go func(key string, delay time.Duration) {
		time.Sleep(delay)
		err := rdb.Del(context.Background(), key).Err()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Unable to delete user from redis.")
		}
	}(uuid, 5*time.Minute)

	ginContext.JSON(http.StatusOK, gin.H{"message": "Confirmation email send."})
}

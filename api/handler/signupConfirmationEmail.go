package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/internal/database/models"
	"github.com/julianopedraca/jubawink/internal/redis"

	log "github.com/sirupsen/logrus"
)

type SendConfirmationEmailResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// @BasePath /api/v1

// @Summary Send Confirmation Email
// @Description Sends a signup confirmation email to the specified user.
// @Tags Email
// @Produce json
// @Success 200 {object} SendConfirmationEmailResponse
// @Failure 400 {object} ErrorResponse
// @Router /email/signup/confirmation [get]
func ConfirmationEmail(ginContext *gin.Context) {
	uuid := ginContext.Param("uuid")
	var rdb = redis.Rdb
	var user models.User
	var jsonMap map[string]interface{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userData, err := rdb.JSONGet(ctx, uuid).Result()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to get the uuid.")
		ginContext.JSON(http.StatusBadGateway, gin.H{"error": "Unable to get the uuid."})
		return
	}
	if userData == "" {
		log.Info("UUID not found")
		ginContext.JSON(http.StatusBadGateway, gin.H{"error": "UUID not found."})
		return
	}

	json.Unmarshal([]byte(userData), &jsonMap)
	user.Email = jsonMap["email"].(string)
	user.UserName = jsonMap["username"].(string)
	user.Password = jsonMap["password"].(string)

	err = user.Save()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to save user to the data base.")
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	err = rdb.Del(context.Background(), uuid).Err()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to delete user from redis.")
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, please try again later."})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Email confirmed and user registered."})
}

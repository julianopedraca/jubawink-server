package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	h "github.com/julianopedraca/jubawink/api/handler"
	"github.com/julianopedraca/jubawink/api/middleware"
	docs "github.com/julianopedraca/jubawink/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RegisterRoutes(server *gin.Engine) {
	server.POST("/login", h.Login)
	server.POST("/signup", h.Signup)
	server.POST("/user/validate", h.ValidateToken)
	server.GET("/info", h.Info)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.GET("/exercise/get/lifting", h.GetExerciseLifting)
	authenticated.POST("/exercise/add/lifting", h.AddExerciseLifting)
	authenticated.GET("/workout/user", h.GetWorkoutsByUserId)
	authenticated.POST("/workout/save", h.SaveWorkout)

	// Initialize swagger
	if os.Getenv("GIN_MODE") == "debug" {
		fmt.Println("Initializing swagger")
		basePath := "/api/v1"
		docs.SwaggerInfo.BasePath = basePath
		swag := server.Group(basePath)
		authenticatedSwag := server.Group(basePath)
		authenticatedSwag.Use(middleware.Authenticate)
		{
			swag.POST("/login", h.Login)
			swag.POST("/signup", h.Signup)
			swag.POST("/user/validate", h.ValidateToken)
			swag.GET("/info", h.Info)

			authenticatedSwag.POST("/exercise/add/lifting", h.AddExerciseLifting)
			authenticatedSwag.GET("/exercise/get/lifting", h.GetExerciseLifting)
			authenticatedSwag.GET("/workout/user", h.GetWorkoutsByUserId)
			authenticatedSwag.POST("/workout/save", h.SaveWorkout)
		}
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

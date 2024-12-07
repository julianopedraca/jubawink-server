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
	server.GET("/email/confirmation/:uuid", h.ConfirmationEmail)
	server.GET("/info", h.Info)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/exercise/add", h.AddExercise)

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
			swag.GET("/email/confirmation", h.ConfirmationEmail)
			swag.GET("/info", h.Info)

			authenticatedSwag.POST("/exercise/add", h.AddExercise)
		}
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

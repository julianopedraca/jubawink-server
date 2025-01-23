package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/api/routes"
	"github.com/julianopedraca/jubawink/internal/database"

	log "github.com/sirupsen/logrus"
)

func main() {
	// connect to postgres
	err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// host := os.Getenv("HOST")
	client := os.Getenv("CLIENT_URL")
	server := gin.Default()

	// Add CORS middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{client}, // Frontend's URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup Security Headers
	server.Use(func(c *gin.Context) {
		// fix for aws so we can prevent server side request forgery (SSRF) https://owasp.org/Top10/pt_BR/A10_2021-Server-Side_Request_Forgery_(SSRF)/
		// if c.Request.Host != host {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
		// 	return
		// }
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	routes.RegisterRoutes(server)

	log.Info("Server running o port 8080.")
	server.Run(":8080")

	database.Db.Close()
}

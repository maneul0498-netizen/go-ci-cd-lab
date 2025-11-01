package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maneul0498-netizen/go-ci-cd-lab/saludo"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			//"message": "¡Hola desde Gin y Render!",
			//"message": "message changued",
			//"message": "changuing message again!!",
			//"message": "changuing message again 3!!",
			//"message": "changuing message again 4!!",
			//"message": "changuing message again 5!!",
			"message": saludo.Message("Manuel"),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}

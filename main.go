package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
			//"message": saludo.Message("Manuel"),
			//"message": "changuing message again 6!!",
			//"message": "changuing message again 7!!",
			//"message": "changuing message again 8!!",
			"message": "changuing message again 9!!",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}

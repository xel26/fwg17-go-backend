package main

import (
	"coffe-shop-be-golang/src/controllers"
	"coffe-shop-be-golang/src/routers"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	}))
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, controllers.ResponseOnly{
			Success: false,
			Message: "Resource not found",
		})
	})
	r.Run("127.0.0.1:8080")
}
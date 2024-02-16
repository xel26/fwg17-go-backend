package main

import (
	"coffe-shop-be-golang/src/routers"
	"coffe-shop-be-golang/src/service"
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
		AllowHeaders: []string{"Content-Type, Authorization"},
	}))
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	r.Static("/uploads", "./uploads")
	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Resource not found",
		})
	})
	r.Run(":8080")
}
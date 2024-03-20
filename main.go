package main

import (
	"coffe-shop-be-golang/src/routers"
	"coffe-shop-be-golang/src/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, Authorization"},
	}))

	// // dimatikan saat build image
	// err := godotenv.Load()
    // if err != nil {
    //     fmt.Println("Error loading .env file")
    // }
	
	r.Static("/uploads", "./uploads")

	routers.Combine(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Resource not found",
		})
	})
	
	// r.Run("127.0.0.1:8080")  //// ==> local computer
	r.Run(":8080") //// ==> build image
}
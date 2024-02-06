package main

import (
	"coffe-shop-be-golang/src/controllers"
	"coffe-shop-be-golang/src/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	routers.Combine(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, controllers.ResponseOnly{
			Success: false,
			Message: "Resource not found",
		})
	})
	r.Run("127.0.0.1:8080")
}
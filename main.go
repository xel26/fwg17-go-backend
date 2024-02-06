package main

import (
	"coffe-shop-be-golang/src/routers"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	routers.Combine(r)
	r.Run("127.0.0.1:8080")
}
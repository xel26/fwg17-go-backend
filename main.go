package main

import (
	"coffe-shop-be-golang/src/routers"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	routers.Combine(r)
	r.Run()
}
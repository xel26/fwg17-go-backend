package middleware

import (
	"coffe-shop-be-golang/src/service"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeToken(c *gin.Context)bool{
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	fmt.Println(authHeader)
	if authHeader == ""{
		return false
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := service.ValidateToken(tokenString)
	if token.Valid{
		claims := token.Claims.(jwt.MapClaims)
		// fmt.Println(claims)
		c.Set("user", claims)
		return true
	}else{
		fmt.Println(err)
		// c.AbortWithStatus(http.StatusUnauthorized)
		return false
	}
}
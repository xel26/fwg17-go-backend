package middleware

import (
	"github.com/gin-gonic/gin"
)

type payload struct{
	email string
	id int
	role string
}

func RecoverPanic()bool{
	if a:= recover(); a != ""{
		return false
	}
	return true
}

func RoleCheck(role string, c *gin.Context) interface{}{
	defer RecoverPanic()
	return c.MustGet("user")
}
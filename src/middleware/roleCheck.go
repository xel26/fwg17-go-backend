package middleware

import (
	"github.com/gin-gonic/gin"
)

type payload struct{
	email string
	id int
	role string
}

func RoleCheck(role string, c *gin.Context) interface{} {
	return c.MustGet("user")
}
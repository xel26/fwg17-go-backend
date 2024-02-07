package middleware

import (
	"coffe-shop-be-golang/src/controllers"
	"coffe-shop-be-golang/src/models"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Auth()(*jwt.GinJWTMiddleware, error){
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "go-backend",
		Key: []byte(os.Getenv("APP_SECRET")),
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			user := data.(*models.User)
			fmt.Println(user)
			return jwt.MapClaims{
				"id": user.Id,
				"role": user.Role,
			}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c) // membaca payload
			return &models.User{
				Id: int(claims["id"].(float64)),
				Role: claims["role"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			form := models.User{}
			err := c.ShouldBind(&form)
		
			found, err := models.FindOneUsersByEmail(form.Email)
		
			if err != nil {
				return nil, err
			}
		
			decode, err := argonize.DecodeHashStr(found.Password)
			
			if err != nil {
				return nil, err
			}
		
			plain := []byte(form.Password)
			if decode.IsValidPassword(plain) {
				return &models.User{
					Id: found.Id,
					Role: found.Role,
				}, nil
			} else {
				return nil, errors.New("invalid_password")
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			user := data.(*models.User)
			if strings.HasPrefix(c.Request.URL.Path, "/admin"){
				if user.Role != "admin"{
					return false
				}
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			if strings.HasPrefix(c.Request.URL.Path, "/login"){
				c.JSON(http.StatusUnauthorized, &controllers.ResponseOnly{
					Success: false,
					Message: "wrong Email or password",
				})
				return
			}

			c.JSON(http.StatusUnauthorized, &controllers.ResponseOnly{
				Success: false,
				Message: "Unauthorized",
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, time time.Time) {
			c.JSON(http.StatusOK, &controllers.Response{
				Success: true,
				Message: "Login success",
				Results: struct{
					Token string `json:"token"`
				}{
					Token: token,
				},
			})
		},
	})

	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
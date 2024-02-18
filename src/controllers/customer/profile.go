package controllers_customer

import (
	"coffe-shop-be-golang/src/lib"
	"coffe-shop-be-golang/src/models"
	"coffe-shop-be-golang/src/service"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/KEINOS/go-argonize"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))


	user, err := models.FindOneUsers(id)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows") {
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "User not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Detail user",
		Results: user,
	})
}

func UpdateProfile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))

	isUserExist, error := models.FindOneUsers(id)
	if error != nil {
		fmt.Println(isUserExist, error)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "no data found",
		})
		return
	}

	data := models.UserForm{}
	err := c.ShouldBind(&data)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	plain := []byte(data.Password)
	hash, err := argonize.Hash(plain)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Password = hash.String()
	data.Id = id

	_, err = c.FormFile("picture")
	if err == nil {
		err = os.Remove("./" + isUserExist.Picture)
		if err != nil{
			fmt.Println(err)
			return
		}

		file, err := lib.Upload(c, "picture", "users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		data.Picture = file
	}else{
		fmt.Println(err)
		data.Picture = ""
	}

	user, err := models.UpdateUser(data)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "User updated successfully",
		Results: user,
	})
}

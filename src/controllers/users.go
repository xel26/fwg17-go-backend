package controllers

import (
	"coffe-shop-be-golang/src/lib"
	"coffe-shop-be-golang/src/models"
	"coffe-shop-be-golang/src/service"
	"fmt"
	"math"
	"os"
	"strings"

	"net/http"
	"strconv"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
)

func ListAllUsers(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllUsers(searchKey, sortBy, order, limit, offset)
	if len(result.Data) == 0 {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "data not found",
		})
		return
	}

	totalPage := int(math.Ceil(float64(result.Count) / float64(limit)))
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = 0
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	PageInfo := service.PageInfo{
		CurrentPage: page,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		Limit:       limit,
		TotalPage:   totalPage,
		TotalData:   result.Count,
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.ResponseList{
		Success:  true,
		Message:  "List all Users",
		PageInfo: PageInfo,
		Results:  result.Data,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.FindOneUsers(id)
	if err != nil {
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

func CreateUser(c *gin.Context) {
	data := models.UserForm{}
	errBind := c.ShouldBind(&data)

	if errBind != nil {
		fmt.Println(errBind)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: errBind.Error(),
		})
		return
	}

	defaultRole := "customer"
	data.Role = &defaultRole

	plain := []byte(data.Password)
	hash, _ := argonize.Hash(plain)
	data.Password = hash.String()

	_, err := c.FormFile("picture")
	if err == nil{
		file, err := lib.Upload(c, "picture", "users")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		data.Picture = file
	}else {
		fmt.Println(err)
		// data.Picture = ""
	}

	user, errDB := models.CreateUser(data)
	if errDB != nil {

		fmt.Println(errDB)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: errDB.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "User created successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.UserForm{}

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	data.Id = id

	isExist, error := models.FindOneUsers(id)
	if error != nil {
		fmt.Println(isExist, error)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "User not found",
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

	_, err = c.FormFile("picture")
	if err == nil {
		_ = os.Remove("./" + isExist.Picture)

		file, err := lib.Upload(c, "picture", "users")
		if err != nil {
			fmt.Println(err)
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
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "User updated successfully",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	isExist, error := models.FindOneUsers(id)
	if error != nil {
		fmt.Println(isExist, error)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "no data found",
		})
		return
	}

	user, err := models.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Delete User Successfully",
		Results: user,
	})
}
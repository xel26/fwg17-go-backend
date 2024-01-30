package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/xel26/fwg17-go-backend/src/models"
)


type pageInfo struct{
	Page int `json:"page"`
}


type responseList struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	PageInfo pageInfo `json:"pageInfo"`
	Results interface{} `json:"results"`
}


type response struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Results interface{} `json:"results"`
}


type User struct{
	Id int `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}


func ListAllUsers(c *gin.Context){
	page, _ := strconv.Atoi(c.Query("page"))

	data, _ := models.GetAllUsers()
	fmt.Println(data)

	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List all Users",
		PageInfo: pageInfo{
			Page: page,
		},
		Results: data,
	})
}


func DetailUser(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Detail user",
		Results: User{
			Id: id,
			Email: "admin@mail.com",
			Password: "1234",
		},
	})
}


func CreateUser(c *gin.Context){
	user := User{}

	c.Bind(&user)

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User created successfully",
		Results: user,
	})
}


func UpdateUser(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{}

	c.Bind(&user)
	user.Id = id

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User updated successfully",
		Results: user,
	})
}


func DeleteUser(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User successfully deleted",
		Results: User{
			Id: id,
			Email: "example@mail.com",
			Password: "1234",
		},
	})
}
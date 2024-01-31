package controllers

import (
	"coffe-shop-be-golang/src/models"
	"log"
	"math"
	"strings"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pageInfo struct {
	CurrentPage int `json:"currentPage"`
	TotalPage int `json:"totalPage"`
	NextPage int `json:"nextPage"`
	PrevPage int `json:"prevPage"`
	Limit int `json:"limit"`
	TotalData int `json:"totalData"`
}

type responseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo pageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type responseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}


func ListAllUsers(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllUsers(searchKey, sortBy, order, limit, offset)


	totalPage := int(math.Ceil(float64(result.Count)/float64(limit)))
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = 0
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	pageInfo := pageInfo{
		CurrentPage: page,
		NextPage: nextPage,
		PrevPage: prevPage,
		Limit: limit,
		TotalPage: totalPage,
		TotalData: result.Count,
	}


	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List all Users",
		PageInfo: pageInfo,
		Results: result.Data,
	})
}


func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.FindOneUsers(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &responseOnly{
				Success: false,
				Message: "User not found",
			})
		return
		}

		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Detail user",
		Results: user,
	})
}


func CreateUser(c *gin.Context) {
	data := models.User{}
	c.Bind(&data)

	user, err := models.CreateUser(data)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User created successfully",
		Results: user,
	})
}


func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.User{}

	c.Bind(&data)
	data.Id = id

	user, err := models.UpdateUser(data)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}


	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "User updated successfully",
		Results: user,
	})
}


func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.DeleteUser(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &responseOnly{
				Success: false,
				Message: "User not found",
			})
		return
		}

		c.JSON(http.StatusInternalServerError, &responseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &response{
		Success: true,
		Message: "Delete User Successfully",
		Results: user,
	})
}
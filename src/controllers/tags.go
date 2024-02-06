package controllers

import (
	"coffe-shop-be-golang/src/models"
	"fmt"
	"math"
	"strings"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func ListAllTags(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllTags(searchKey, sortBy, order, limit, offset)
	if len(result.Data) == 0 {
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "data not found",
		})
		return
	}


	totalPage := int(math.Ceil(float64(result.Count)/float64(limit)))
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = 0
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	PageInfo := PageInfo{
		CurrentPage: page,
		NextPage: nextPage,
		PrevPage: prevPage,
		Limit: limit,
		TotalPage: totalPage,
		TotalData: result.Count,
	}


	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &ResponseList{
		Success: true,
		Message: "List all tags",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tag, err := models.FindOneTags(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &ResponseOnly{
				Success: false,
				Message: "Tags not found",
			})
		return
		}

		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Detail tags",
		Results: tag,
	})
}


func CreateTags(c *gin.Context) {
	data := models.Tags{}
	c.ShouldBind(&data)

	tags, err := models.CreateTags(data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Tags created successfully",
		Results: tags,
	})
}


func UpdateTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.Tags{}

	c.ShouldBind(&data)
	data.Id = id

	isExist, err := models.FindOneTags(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Tags not found",
		})
	return
	}

	tags, err := models.UpdateTags(data)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &ResponseOnly{
				Success: false,
				Message: "Tags not found",
			})
		return
		}
		
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}


	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Tags updated successfully",
		Results: tags,
	})
}


func DeleteTags(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneTags(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Tags not found",
		})
	return
	}

	tags, err := models.DeleteTags(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Delete tags successfully",
		Results: tags,
	})
}
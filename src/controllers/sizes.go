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


func ListAllSizes(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllSizes(searchKey, sortBy, order, limit, offset)
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
		Message: "List all sizes",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailSizes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	size, err := models.FindOneSizes(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &ResponseOnly{
				Success: false,
				Message: "Sizes not found",
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
		Message: "Detail size",
		Results: size,
	})
}


func CreateSize(c *gin.Context) {
	data := models.SizesForm{}
	c.ShouldBind(&data)

	size, err := models.CreateSizes(data)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "pq: duplicate key"){
			c.JSON(http.StatusBadRequest, &ResponseOnly{
				Success: false,
				Message: "duplicate size name",
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
		Message: "Sizes created successfully",
		Results: size,
	})
}


func UpdateSizes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.SizesForm{}

	c.ShouldBind(&data)
	data.Id = id

	isExist, err := models.FindOneSizes(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Size not found",
		})
	return
	}

	size, err := models.UpdateSizes(data)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &ResponseOnly{
				Success: false,
				Message: "Sizes not found",
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
		Message: "Sizes updated successfully",
		Results: size,
	})
}


func DeleteSizes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneSizes(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Size not found",
		})
	return
	}

	sizes, err := models.DeleteSizes(id)
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
		Message: "Delete sizes successfully",
		Results: sizes,
	})
}
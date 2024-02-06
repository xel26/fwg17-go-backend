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


func ListAllPromo(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllPromo(searchKey, sortBy, order, limit, offset)
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
		Message: "List all promo",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailPromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	promo, err := models.FindOnePromo(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &ResponseOnly{
				Success: false,
				Message: "Promo not found",
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
		Message: "Detail promo",
		Results: promo,
	})
}


func CreatePromo(c *gin.Context) {
	data := models.PromoForm{}
	c.ShouldBind(&data)

	promo, err := models.CreatePromo(data)
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
		Message: "Promo created successfully",
		Results: promo,
	})
}


func UpdatePromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.PromoForm{}

	err := c.ShouldBind(&data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}
	data.Id = id

	isExist, err := models.FindOnePromo(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Promo not found",
		})
	return
	}

	fmt.Println(data)
	promo, err := models.UpdatePromo(data)
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
		Message: "Promo updated successfully",
		Results: promo,
	})
}


func DeletePromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOnePromo(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Promo not found",
		})
	return
	}

	promo, err := models.DeletePromo(id)
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
		Message: "Delete promo successfully",
		Results: promo,
	})
}
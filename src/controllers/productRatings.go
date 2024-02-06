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


func ListAllProductRatings(c *gin.Context) {
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllProductRatings(sortBy, order, limit, offset)
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
		Message: "List all product ratings",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailProductRatings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pr, err := models.FindOneProductRatings(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusNotFound, &ResponseOnly{
				Success: false,
				Message: "Product ratings not found",
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
		Message: "Detail product ratings",
		Results: pr,
	})
}


func CreateProductRatings(c *gin.Context) {
	data := models.ProductRatings{}
	c.ShouldBind(&data)

	_, err := models.FindOneProducts(data.ProductId)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "product id not found",
		})
		return
	}

	_, err = models.FindOneUsers(data.UserId)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "user id not found",
		})
		return
	}

	dataForm := models.PRForm{}
	c.ShouldBind(&dataForm)
	fmt.Println(dataForm)

	pr, err := models.CreateProductRatings(dataForm)
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
		Message: "Product ratings created successfully",
		Results: pr,
	})
}


func UpdatePrductRatings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.ProductRatings{}

	c.ShouldBind(&data)

	_, err := models.FindOneProducts(data.ProductId)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "product id not found",
		})
		return
	}

	_, err = models.FindOneUsers(data.UserId)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "user id not found",
		})
		return
	}

	dataForm := models.PRForm{}
	c.ShouldBind(&dataForm)
	dataForm.Id = id
	fmt.Println(dataForm)

	pr, err := models.UpdateProductRatings(dataForm)
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
		Message: "Product ratings updated successfully",
		Results: pr,
	})
}


func DeleteProductRatings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneProductRatings(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Product ratings not found",
		})
	return
	}

	pr, err := models.DeleteProductRatings(id)
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
		Message: "Delete product successfully",
		Results: pr,
	})
}
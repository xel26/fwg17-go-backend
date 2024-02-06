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


func ListAllProductVariants(c *gin.Context) {
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllProductVariants(sortBy, order, limit, offset)
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
		Message: "List all product variants",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailProductVariant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pv, err := models.FindOneProductVariants(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &ResponseOnly{
				Success: false,
				Message: "Product variants not found",
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
		Message: "Detail product variants",
		Results: pv,
	})
}


func CreateProductVariants(c *gin.Context) {
	data := models.ProductVariants{}
	c.ShouldBind(&data)

	_, err := models.FindOneProducts(data.ProductId)
	if err != nil{
		fmt.Println(err, data)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "product id not found",
		})
		return
	}

	_, err = models.FindOneVariants(data.VariantId)
	if err != nil{
		fmt.Println(err, data)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Variant id not found",
		})
		return
	}

	pv, err := models.CreateProductVariants(data)
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
		Message: "Product variants created successfully",
		Results: pv,
	})
}


func UpdateProductVariants(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.ProductVariants{}

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

	isExist, err := models.FindOneProductVariants(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Product variants not found",
		})
	return
	}

	_, err = models.FindOneVariants(data.VariantId)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Variant id not found",
		})
		return
	}

	data.Id = id

	pv, err := models.UpdateProductVariants(data)
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
		Message: "Product variant updated successfully",
		Results: pv,
	})
}


func DeleteProductVariants(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneProductVariants(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Product varinants not found",
		})
	return
	}


	pv, err := models.DeleteProductVariants(id)
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
		Message: "Delete product variants successfully",
		Results: pv,
	})
}
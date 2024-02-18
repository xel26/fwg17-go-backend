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

	"github.com/gin-gonic/gin"
)


func ListAllProducts(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	category := c.DefaultQuery("category", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllProducts(searchKey, category, sortBy, order, limit, offset)
	if len(result.Data) == 0 {
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
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

	PageInfo := service.PageInfo{
		CurrentPage: page,
		NextPage: nextPage,
		PrevPage: prevPage,
		Limit: limit,
		TotalPage: totalPage,
		TotalData: result.Count,
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
		Success: true,
		Message: "List all products",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailProducts(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.FindOneProducts(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusNotFound, &service.ResponseOnly{
				Success: false,
				Message: "Product not found",
			})
		return
		}

		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Detail product",
		Results: product,
	})
}


func CreateProducts(c *gin.Context) {
	data := models.ProductForm{}
	errBind := c.ShouldBind(&data)
	if errBind != nil {
		fmt.Println(errBind)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}


	_, err := c.FormFile("image")
	if err == nil {
		file, err := lib.Upload(c, "image", "products")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		data.Image = file
	}else{
		fmt.Println(err)
		data.Image = ""
	}



	product, errDB := models.CreateProducts(data)
	if errDB != nil {
		fmt.Println(errDB)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Products created successfully",
		Results: product,
	})
}


func UpdatePrducts(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.ProductForm{}

	errBind := c.ShouldBind(&data)
	if errBind != nil{
		fmt.Println(errBind)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: errBind.Error(),
		})
	return
	}

	data.Id = id


	isExist, err := models.FindOneProducts(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Product not found",
		})
	return
	}


	_, err = c.FormFile("image")
	if err == nil {
		_ = os.Remove("./" + isExist.Image)

		file, err := lib.Upload(c, "image", "products")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		data.Image = file
	}else{
		fmt.Println(err)
		data.Image = ""
	}


	product, err := models.UpdateProduct(data)
	if err != nil {
		fmt.Println(err, product)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Product updated successfully",
		Results: product,
	})
}


func DeleteProducts(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	isExist, err := models.FindOneProducts(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "Product not found",
		})
	return
	}
	_ = os.Remove("./" + isExist.Image)


	product, err := models.DeleteProduct(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "Delete product successfully",
		Results: product,
	})
}
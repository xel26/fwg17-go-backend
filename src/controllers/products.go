package controllers

import (
	"coffe-shop-be-golang/src/middleware"
	"coffe-shop-be-golang/src/models"
	"fmt"
	"log"
	"math"
	"strings"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func ListAllProducts(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}


	searchKey := c.DefaultQuery("searchKey", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllProducts(searchKey, sortBy, order, limit, offset)
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
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &ResponseList{
		Success: true,
		Message: "List all products",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailProducts(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}



	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.FindOneProducts(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusNotFound, &ResponseOnly{
				Success: false,
				Message: "Product not found",
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
		Message: "Detail product",
		Results: product,
	})
}


func CreateProducts(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}


	data := models.Product{}
	c.ShouldBind(&data)

	product, err := models.CreateProducts(data)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Products created successfully",
		Results: product,
	})
}


func UpdatePrducts(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}


	id, _ := strconv.Atoi(c.Param("id"))
	data := models.Product{}

	c.ShouldBind(&data)
	data.Id = id


	isExist, err := models.FindOneProducts(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Product not found",
		})
	return
	}


	product, err := models.UpdateProduct(data)
	if err != nil {
		fmt.Println(err, product)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal server error",
		})
		return
	}


	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Product updated successfully",
		Results: product,
	})
}


func DeleteProducts(c *gin.Context) {
	isAuthorize := middleware.AuthorizeToken(c)
	claims := middleware.RoleCheck("admin", c)
	fmt.Println("claims", claims)


	if isAuthorize == false || claims == false{
		c.JSON(http.StatusUnauthorized, &ResponseOnly{
			Success: false,
			Message: "Unauthorize",
		})
		return
	}

	
	id, _ := strconv.Atoi(c.Param("id"))
	isExist, err := models.FindOneProducts(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Product not found",
		})
	return
	}


	product, err := models.DeleteProduct(id)
	if err != nil {
		fmt.Println(err, product)
		c.JSON(http.StatusInternalServerError, &ResponseOnly{
			Success: false,
			Message: "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Success: true,
		Message: "Delete product successfully",
		Results: product,
	})
}
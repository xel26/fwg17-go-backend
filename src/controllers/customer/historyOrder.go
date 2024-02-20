package controllers_customer

import (
	"coffe-shop-be-golang/src/models"
	"coffe-shop-be-golang/src/service"
	"fmt"
	"math"
	"strings"

	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)


func ListAllOrders(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := int(claims["id"].(float64))
	
	status := c.DefaultQuery("status", "")
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "DESC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	offset := (page - 1) * limit

	result, err := models.FindAllOrdersByUserId(status, userId, sortBy, order, limit, offset)
	if len(result.Data) == 0 {
		fmt.Println(err)
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
		Message: "List all orders",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailOrder(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := int(claims["id"].(float64))

	id, _ := strconv.Atoi(c.Param("id"))
	orders, err := models.FindOneOrderByUserId(id, userId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "Order not found",
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
		Message: "Detail order",
		Results: orders,
	})
}


func ListOrderProducts(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := int(claims["id"].(float64))
	
	orderId, _ := strconv.Atoi(c.Query("orderId"))
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	result, err := models.GetOrderProducts(orderId, userId, sortBy, order)
	if err == nil && len(result.Data) == 0 {
		c.JSON(http.StatusNotFound, &service.ResponseOnly{
			Success: false,
			Message: "data not found",
		})
		return
	}


	totalPage := 1
	nextPage := 0
	prevPage := 0

	PageInfo := service.PageInfo{
		CurrentPage: page,
		NextPage: nextPage,
		PrevPage: prevPage,
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
		Message: "List all order products",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}
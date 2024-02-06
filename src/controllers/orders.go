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


func ListAllOrders(c *gin.Context) {
	sortBy := c.DefaultQuery("sortBy", "id")
	order := c.DefaultQuery("order", "ASC")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	result, err := models.FindAllOrders(sortBy, order, limit, offset)
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
		Message: "List all orders",
		PageInfo: PageInfo,
		Results: result.Data,
	})
}


func DetailOrders(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	orders, err := models.FindOneOrders(id)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &ResponseOnly{
				Success: false,
				Message: "Order not found",
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
		Message: "Detail orders",
		Results: orders,
	})
}


func CreateOrders(c *gin.Context) {
	data := models.OrderForm{}
	err := c.ShouldBind(&data)

	order, err := models.CreateOrders(data)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "pq: duplicate key"){
			c.JSON(http.StatusBadRequest, &ResponseOnly{
				Success: false,
				Message: "duplicate key for orderNumber",
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
		Message: "Order created successfully",
		Results: order,
	})
}


func UpdateOrders(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.OrderForm{}

	c.ShouldBind(&data)
	data.Id = id

	isExist, err := models.FindOneOrders(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Order not found",
		})
	return
	}

	orders, err := models.UpdateOrders(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "pq: duplicate key"){
			c.JSON(http.StatusBadRequest, &ResponseOnly{
				Success: false,
				Message: "duplicate key for orderNumber",
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
		Message: "Order updated successfully",
		Results: orders,
	})
}


func DeleteOrders(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	isExist, err := models.FindOneOrders(id)
	if err != nil{
		fmt.Println(isExist, err)
		c.JSON(http.StatusNotFound, &ResponseOnly{
			Success: false,
			Message: "Order not found",
		})
	return
	}

	orders, err := models.DeleteOrders(id)
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
		Message: "Delete Order successfully",
		Results: orders,
	})
}
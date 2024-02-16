package controllers_customer

import (
	"coffe-shop-be-golang/src/lib"
	"coffe-shop-be-golang/src/models"
	"coffe-shop-be-golang/src/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = lib.DB

func GetPriceSize(c *gin.Context) {
	size := c.Query("size")
	dataSize, err := models.GetOneSize(size)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "Size not found",
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
		Message: "Data size",
		Results: dataSize,
	})
}


func GetPriceVariant(c *gin.Context) {
	name := c.Query("name")
	dataVariant, err := models.GetOneVariant(name)
	if err != nil {
		fmt.Println(err)
		if strings.HasPrefix(err.Error(), "sql: no rows"){
			c.JSON(http.StatusInternalServerError, &service.ResponseOnly{
				Success: false,
				Message: "Variants not found",
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
		Message: "Detail variants",
		Results: dataVariant,
	})
}


func Checkout(c *gin.Context) {
	tx, err := db.BeginTx(c, nil)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	defer tx.Rollback()


	claims := jwt.ExtractClaims(c)
	userId := int(claims["id"].(float64))

	dataOrder := models.CheckoutForm{}
	err = c.ShouldBind(&dataOrder)
	if err != nil {
		fmt.Println("test", err)
		tx.Rollback()
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	dataOrder.OrderNumber = lib.RandomNumberStr(9)
	dataOrder.UserId = userId
	dataOrder.Status = "On Progress"
	dataOrder.DeliveryFee = 5000



	if c.PostForm("deliveryAddress") == ""{
		user, err := models.GetAddress(userId)
		fmt.Println("test", err, user)
		if err != nil {
			fmt.Println(err, user)
			tx.Rollback()
			c.JSON(http.StatusBadRequest, &service.ResponseOnly{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		dataOrder.DeliveryAddress = user.Address
	}

	if c.PostForm("fullName") == ""{
		user, err := models.GetFullName(userId)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return
		}

		dataOrder.FullName = user.FullName
	}

	if c.PostForm("email") == ""{
		user, err := models.GetEmail(userId)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return
		}

		dataOrder.Email = user.Email
	}

	order, err := models.InsertOrder(dataOrder)
	if err != nil {
		fmt.Println("error aja", err, order)
		tx.Rollback()
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	dataOrder.OrderId = order.Id
	fmt.Println("data order", dataOrder.OrderId)


	data := models.OD{}
	data.OrderId = order.Id

	productId := strings.Split(dataOrder.ProductId, ",")
	size := strings.Split(dataOrder.SizeProduct, ",")
	variant := strings.Split(dataOrder.VariantProduct, ",")
	quantityProduct := strings.Split(dataOrder.QuantityProduct, ",")

	for i := 0; i < len(productId); i++{
		sizeId, _ := models.GetOneSize(size[i])
		variantId, _ := models.GetOneVariant(variant[i])

		data.ProductId, _ = strconv.Atoi(productId[i])
		data.Quantity, _ = strconv.Atoi(quantityProduct[i])
		data.SizeId = sizeId.Id
		data.VariantId = variantId.Id


		orderDetails, _ := models.CreateOD(data)
		models.CountSubtotal(orderDetails.Id)
	}

	_, _ = models.CountTotal(order.Id)
	_, _ = models.CountTax(order.Id)

	order, err = models.CountTotalTransaction(order.Id)

	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = tx.Commit()
	
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		c.JSON(http.StatusBadRequest, &service.ResponseOnly{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "create order successfully",
		Results: order,
	})

}

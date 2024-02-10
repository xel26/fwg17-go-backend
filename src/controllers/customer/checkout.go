package controllers_customer

import (
	"coffe-shop-be-golang/src/lib"
	"coffe-shop-be-golang/src/models"
	"coffe-shop-be-golang/src/service"
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = lib.DB

func Checkout(c *gin.Context){
	claims := jwt.ExtractClaims(c)
	userId := int(claims["id"].(float64))

	dataOrder := models.OrderForm{}
	err := c.ShouldBind(&dataOrder) 
	if err != nil{
		fmt.Println(err)
		return
	}

	dataOrder.UserId = &userId
	status := "On Progress"
	deliveryFee := 5000
	dataOrder.Status = &status
	dataOrder.DeliveryFee = &deliveryFee

	fmt.Println(dataOrder)
	// order, err := models.CreateOrders(data)

	tx, err := db.BeginTx(c, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback()


	err = tx.Commit()
	if err != nil{
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, &service.Response{
		Success: true,
		Message: "create order successfully",
		Results: dataOrder,
	})

}
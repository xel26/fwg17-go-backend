package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbOD *sqlx.DB = lib.DB

type OrderDetails struct {
	Id        int          `dbOD:"id" json:"id"`
	OrderDetailsId int          `dbOD:"productId" json:"productId" form:"productId"`
	SizeId    int          `dbOD:"sizeId" json:"sizeId" form:"sizeId"`
	VariantId int          `dbOD:"variantId" json:"variantId" form:"variantId"`
	Quantity  int          `dbOD:"quantity" json:"quantity" form:"quantity"`
	OrderId   int          `dbOD:"orderId" json:"orderId" form:"orderId"`
	Subtotal  int          `dbOD:"subtotal" json:"subtotal" form:"subtotal"`
	CreatedAt time.Time    `dbOD:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `dbOD:"updatedAt" json:"updatedAt"`
}

type InfoOD struct {
	Data  []OrderDetails
	Count int
}



func FindAllOrderDetails(sortBy string, order string, limit int, offset int) (InfoOD, error) {
	sql := `
	SELECT * FROM "orderDetails" 
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "orderDetails"
	`

	result := InfoOD{}
	data := []OrderDetails{}
	err := dbOD.Select(&data, sql, limit, offset)
	result.Data = data
	
	row := dbOD.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}



func FindOneOrderDetails(id int) (OrderDetails, error) {
	sql := `SELECT * FROM "orderDetails" WHERE id = $1`
	data := OrderDetails{}
	err := dbOD.Get(&data, sql, id)
	return data, err
}



func CreateOrderDetails(data OrderDetails) (OrderDetails, error) {
	sql := `INSERT INTO "orderDetails" ("productId", "sizeId", "variantId", "quantity", "orderId", "subtotal") 
	VALUES
	(:productId, :sizeId, :variantId, :quantity, :orderId, :subtotal)
	RETURNING *
	`
	result := OrderDetails{}
	rows, err := dbOD.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}
	
	for rows.Next(){
		rows.StructScan(&result)
	}
	
	return result, err
}




func UpdateOrderDetails(data OrderDetails) (OrderDetails, error) {
	sql := `UPDATE "orderDetails" SET
	"productId"=COALESCE(NULLIF(:productId, ''),"productId"),
	"sizeId"=COALESCE(NULLIF(:sizeId, ''),"sizeId"),
	"variantId"=COALESCE(NULLIF(:variantId, ''),"variantId"),
	"quantity"=COALESCE(NULLIF(:quantity, ''),"quantity"),
	"orderId"=COALESCE(NULLIF(:orderId, ''),"orderId"),
	"subtotal"=COALESCE(NULLIF(:subtotal, ''),"subtotal"),
	WHERE id=:id
	RETURNING *
	`
	result := OrderDetails{}
	rows, err := dbOD.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next(){
		rows.StructScan(&result)
	}

	return result, err
}



func DeleteOrderDetails(id int) (OrderDetails, error) {
	sql := `DELETE FROM "orderDetails" WHERE id = $1 RETURNING *`
	data := OrderDetails{}
	err := dbOD.Get(&data, sql, id)
	return data, err
}
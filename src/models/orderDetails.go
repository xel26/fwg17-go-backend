package models

import (
	"database/sql"
	"time"
)

type OrderDetails struct {
	Id        int          `db:"id" json:"id"`
	ProductId int          `db:"productId" json:"productId" form:"productId" binding:"required,numeric"`
	SizeId    int          `db:"sizeId" json:"sizeId" form:"sizeId" binding:"required,numeric"`
	VariantId int          `db:"variantId" json:"variantId" form:"variantId" binding:"required,numeric"`
	Quantity  int          `db:"quantity" json:"quantity" form:"quantity" binding:"required,numeric"`
	OrderId   int          `db:"orderId" json:"orderId" form:"orderId" binding:"required,numeric"`
	Subtotal  int          `db:"subtotal" json:"subtotal" form:"subtotal" binding:"required,numeric"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoOD struct {
	Data  []OrderDetails
	Count int
}

func FindAllOrderDetails(sortBy string, order string, limit int, offset int) (InfoOD, error) {
	sql := `
	SELECT * FROM "orderDetails" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "orderDetails"
	`

	result := InfoOD{}
	data := []OrderDetails{}
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneOrderDetails(id int) (OrderDetails, error) {
	sql := `SELECT * FROM "orderDetails" WHERE id = $1`
	data := OrderDetails{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateOrderDetails(data OrderDetails) (OrderDetails, error) {
	sql := `INSERT INTO "orderDetails" ("productId", "sizeId", "variantId", "quantity", "orderId", "subtotal") 
	VALUES
	(:productId, :sizeId, :variantId, :quantity, :orderId, :subtotal)
	RETURNING *
	`
	result := OrderDetails{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateOrderDetails(data OrderDetails) (OrderDetails, error) {
	sql := `UPDATE "orderDetails" SET
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"sizeId"=COALESCE(NULLIF(:sizeId, 0),"sizeId"),
	"variantId"=COALESCE(NULLIF(:variantId, 0),"variantId"),
	"quantity"=COALESCE(NULLIF(:quantity, 0),"quantity"),
	"orderId"=COALESCE(NULLIF(:orderId, 0),"orderId"),
	"subtotal"=COALESCE(NULLIF(:subtotal, 0),"subtotal"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := OrderDetails{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteOrderDetails(id int) (OrderDetails, error) {
	sql := `DELETE FROM "orderDetails" WHERE id = $1 RETURNING *`
	data := OrderDetails{}
	err := db.Get(&data, sql, id)
	return data, err
}

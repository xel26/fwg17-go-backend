package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbO *sqlx.DB = lib.DB

type Orders struct {
	Id               int            `dbO:"id" json:"id"`
	UserId           int            `dbO:"userId" json:"userId" form:"userId"`
	OrderNumber      string         `dbO:"orderNumber" json:"orderNumber" form:"orderNumber"`
	PromoId          sql.NullInt64  `dbO:"promoId" json:"promoId" form:"promoId"`
	Total            int            `dbO:"total" json:"total" form:"total"`
	Tax              sql.NullInt64  `dbO:"tax" json:"tax" form:"tax"`
	DeliveryAddress  sql.NullString `dbO:"deliveryAddress" json:"deliveryAddress" form:"deliveryAddress"`
	FullName         sql.NullString `dbO:"fullName" json:"fullName" form:"fullName"`
	Email            sql.NullString `dbO:"email" json:"email" form:"email"`
	PriceCut         sql.NullInt64  `dbO:"priceCut" json:"priceCut" form:"priceCut"`
	Subtotal         int            `dbO:"subtotal" json:"subtotal" form:"subtotal"`
	Status           string         `dbO:"status" json:"status" form:"status"`
	DeliveryFee      int            `dbO:"deliveryFee" json:"deliveryFee" form:"deliveryFee"`
	DeliveryShipping string         `dbO:"deliveryShipping" json:"deliveryShipping" form:"deliveryShipping"`
	CreatedAt        time.Time      `dbO:"createdAt" json:"createdAt"`
	UpdatedAt        sql.NullTime   `dbO:"updatedAt" json:"updatedAt"`
}

type InfoO struct {
	Data  []Orders
	Count int
}

func FindAllOrders(sortBy string, order string, limit int, offset int) (InfoO, error) {
	sql := `
	SELECT * FROM "orders" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "orders"
	`

	result := InfoO{}
	data := []Orders{}
	err := dbO.Select(&data, sql, limit, offset)
	result.Data = data

	row := dbO.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneOrders(id int) (Orders, error) {
	sql := `SELECT * FROM "orders" WHERE id = $1`
	data := Orders{}
	err := dbO.Get(&data, sql, id)
	return data, err
}

func CreateOrders(data Orders) (Orders, error) {
	sql := `INSERT INTO "orders" ("name", "description", "basePrice", "image", "discount", "isRecommended", "tagId") 
	VALUES
	(:name, :description, :basePrice, :image, :discount, :isRecommended, :tagId)
	RETURNING *
	`
	result := Orders{}
	rows, err := dbO.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateOrders(data Orders) (Orders, error) {
	sql := `UPDATE "orders" SET
	"userId"=COALESCE(NULLIF(:userId, ''),"userId"),
	"orderNumber"=COALESCE(NULLIF(:orderNumber, ''),"orderNumber"),
	"promoId"=COALESCE(NULLIF(:promoId, ''),"promoId"),
	"total"=COALESCE(NULLIF(:total, ''),"total"),
	"deliveryAddress"=COALESCE(NULLIF(:deliveryAddress, ''),"deliveryAddress"),
	"fullName"=COALESCE(NULLIF(:fullName, ''),"fullName"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"priceCut"=COALESCE(NULLIF(:priceCut, ''),"priceCut"),
	"subtotal"=COALESCE(NULLIF(:subtotal, ''),"subtotal"),
	"status"=COALESCE(NULLIF(:status, ''),"status"),
	"deliveryFee"=COALESCE(NULLIF(:deliveryFee, ''),"deliveryFee"),
	"deliveryShipping"=COALESCE(NULLIF(:deliveryShipping, ''),"deliveryShipping"),
	WHERE id=:id
	RETURNING *
	`
	result := Orders{}
	rows, err := dbO.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteOrders(id int) (Orders, error) {
	sql := `DELETE FROM "orders" WHERE id = $1 RETURNING *`
	data := Orders{}
	err := dbO.Get(&data, sql, id)
	return data, err
}

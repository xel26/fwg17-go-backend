package models

import (
	"database/sql"
	"time"
)

type Orders struct {
	Id               int            `db:"id" json:"id"`
	UserId           int            `db:"userId" json:"userId" form:"userId"`
	OrderNumber      string         `db:"orderNumber" json:"orderNumber" form:"orderNumber"`
	PromoId          sql.NullInt64  `db:"promoId" json:"promoId" form:"promoId"`
	Total            int            `db:"total" json:"total" form:"total"`
	Tax              sql.NullInt64  `db:"tax" json:"tax" form:"tax"`
	DeliveryAddress  sql.NullString `db:"deliveryAddress" json:"deliveryAddress" form:"deliveryAddress"`
	FullName         sql.NullString `db:"fullName" json:"fullName" form:"fullName"`
	Email            sql.NullString `db:"email" json:"email" form:"email"`
	PriceCut         sql.NullInt64  `db:"priceCut" json:"priceCut" form:"priceCut"`
	Subtotal         int            `db:"subtotal" json:"subtotal" form:"subtotal"`
	Status           string         `db:"status" json:"status" form:"status"`
	DeliveryFee      int            `db:"deliveryFee" json:"deliveryFee" form:"deliveryFee"`
	DeliveryShipping string         `db:"deliveryShipping" json:"deliveryShipping" form:"deliveryShipping"`
	CreatedAt        time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt        sql.NullTime   `db:"updatedAt" json:"updatedAt"`
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
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneOrders(id int) (Orders, error) {
	sql := `SELECT * FROM "orders" WHERE id = $1`
	data := Orders{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateOrders(data Orders) (Orders, error) {
	sql := `INSERT INTO "orders" ("name", "description", "basePrice", "image", "discount", "isRecommended", "tagId") 
	VALUES
	(:name, :description, :basePrice, :image, :discount, :isRecommended, :tagId)
	RETURNING *
	`
	result := Orders{}
	rows, err := db.NamedQuery(sql, data)
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
	rows, err := db.NamedQuery(sql, data)
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
	err := db.Get(&data, sql, id)
	return data, err
}

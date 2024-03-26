package models

import (
	"time"

	"github.com/LukaGiorgadze/gonull"
	"github.com/lib/pq"
)

type Order struct {
	Id               int            `db:"id" json:"id"`
	UserId           int            `db:"userId" json:"userId" form:"userId"`
	OrderNumber      gonull.Nullable[string] `db:"orderNumber" json:"orderNumber" form:"orderNumber"`
	PromoId          gonull.Nullable[int]  `db:"promoId" json:"promoId" form:"promoId"`
	Total            gonull.Nullable[int]  `db:"total" json:"total" form:"total"`
	Tax              gonull.Nullable[int]  `db:"tax" json:"tax" form:"tax"`
	DeliveryAddress  string         `db:"deliveryAddress" json:"deliveryAddress" form:"deliveryAddress"`
	FullName         gonull.Nullable[string] `db:"fullName" json:"fullName" form:"fullName"`
	Email            gonull.Nullable[string] `db:"email" json:"email" form:"email"`
	PriceCut         gonull.Nullable[int]  `db:"priceCut" json:"priceCut" form:"priceCut"`
	Subtotal         gonull.Nullable[int]  `db:"subtotal" json:"subtotal" form:"subtotal"`
	Status           string         `db:"status" json:"status" form:"status"`
	DeliveryFee      gonull.Nullable[int]  `db:"deliveryFee" json:"deliveryFee" form:"deliveryFee"`
	DeliveryShipping gonull.Nullable[string] `db:"deliveryShipping" json:"deliveryShipping" form:"deliveryShipping"`
	ProductsImage pq.StringArray `db:"productsImage" json:"productsImage"`
	CreatedAt        time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt        gonull.Nullable[time.Time]   `db:"updatedAt" json:"updatedAt"`
}

type OrderForm struct {
	Id               int          `db:"id" json:"id"`
	UserId           *int         `db:"userId" json:"userId" form:"userId"`
	OrderNumber      *string      `db:"orderNumber" json:"orderNumber" form:"orderNumber"`
	PromoId          *int         `db:"promoId" json:"promoId" form:"promoId"`
	Total            *int         `db:"total" json:"total" form:"total"`
	Tax              *int         `db:"tax" json:"tax" form:"tax"`
	DeliveryAddress  *string      `db:"deliveryAddress" json:"deliveryAddress" form:"deliveryAddress"`
	FullName         *string      `db:"fullName" json:"fullName" form:"fullName"`
	Email            *string      `db:"email" json:"email" form:"email"`
	PriceCut         *int         `db:"priceCut" json:"priceCut" form:"priceCut"`
	Subtotal         *int         `db:"subtotal" json:"subtotal" form:"subtotal"`
	Status           *string      `db:"status" json:"status" form:"status"`
	DeliveryFee      *int         `db:"deliveryFee" json:"deliveryFee" form:"deliveryFee"`
	DeliveryShipping *string      `db:"deliveryShipping" json:"deliveryShipping" form:"deliveryShipping"`
	CreatedAt        time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt        gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}

type CheckoutForm struct {
	Id               int          `db:"id" json:"id"`
	UserId           int          `db:"userId" json:"userId"`
	OrderNumber      string       `db:"orderNumber" json:"orderNumber"`
	DeliveryAddress  *string      `db:"deliveryAddress" json:"deliveryAddress" form:"deliveryAddress"`
	FullName         *string      `db:"fullName" json:"fullName" form:"fullName"`
	Email            *string      `db:"email" json:"email" form:"email"`
	DeliveryFee      int          `db:"deliveryFee" json:"deliveryFee"`
	DeliveryShipping *string      `db:"deliveryShipping" json:"deliveryShipping" form:"deliveryShipping" binding:"eq=Dine In|eq=Pick Up|eq=Door Delivery"`
	Status           string       `db:"status" json:"status"`
	OrderId          int          `db:"orderId" json:"orderId"`
	ProductId        string       `db:"productId" json:"productId" form:"productId"`
	SizeProduct      string       `db:"sizeProduct" json:"sizeProduct" form:"sizeProduct"`
	VariantProduct   string       `db:"variantProduct" json:"variantProduct" form:"variantProduct"`
	QuantityProduct  string       `db:"quantityProduct" json:"quantityProduct" form:"quantityProduct"`
	Tax              *int         `db:"tax" json:"tax"`
	Total            *int         `db:"total" json:"total"`
	Subtotal         *int         `db:"subtotal" json:"subtotal"`
	CreatedAt        time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt        gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}

type InfoO struct {
	Data  []Order
	Count int
}

func FindAllOrders(deliveryShipping string, sortBy string, order string, limit int, offset int, status string) (InfoO, error) {
	sql := `
	SELECT "o".*,
    array_agg(DISTINCT "p"."image") "productsImage"
    FROM "orders" "o"
    JOIN "orderDetails" "od" ON ("od"."orderId" = "o"."id")
    JOIN "products" "p" ON ("p"."id" = "od"."productId")
	WHERE "deliveryShipping" ILIKE $1 OR "status" ILIKE $4
	GROUP BY "o"."id"
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "orders"
	WHERE "deliveryShipping" ILIKE $1 OR "status" ILIKE $2
	`

	result := InfoO{}
	data := []Order{}
	error := db.Select(&data, sql, "%"+deliveryShipping+"%", limit, offset, status)
	if error != nil {
		return result, error
	}

	result.Data = data

	row := db.QueryRow(sqlCount, "%"+deliveryShipping+"%", status)
	err := row.Scan(&result.Count)

	return result, err
}

func FindOneOrders(id int) (Order, error) {
	sql := `SELECT * FROM "orders" WHERE id = $1`
	data := Order{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneOrderByUserId(id int, userId int) (Order, error) {
	sql := `SELECT * FROM "orders" WHERE "id" = $1 AND "userId" = $2`
	data := Order{}
	err := db.Get(&data, sql, id, userId)
	return data, err
}

func CreateOrders(data OrderForm) (OrderForm, error) {
	sql := `INSERT INTO "orders" ("userId", "orderNumber", "promoId", "total", "deliveryAddress", "fullName", "email", "priceCut", "subtotal", "status", "deliveryFee", "deliveryShipping", "tax") 
	VALUES
	(:userId, :orderNumber, :promoId, :total, :deliveryAddress, :fullName, :email, :priceCut, :subtotal, :status, :deliveryFee, :deliveryShipping, :tax)
	RETURNING *
	`
	result := OrderForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func InsertOrder(data CheckoutForm) (OrderForm, error) {
	sql := `
	INSERT INTO "orders" ("userId", "orderNumber", "deliveryAddress", "fullName", "email", "status", "deliveryFee", "deliveryShipping") 
	VALUES
	(:userId, :orderNumber, :deliveryAddress, :fullName, :email, :status, :deliveryFee, :deliveryShipping)
	RETURNING *
	`
	result := OrderForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateOrders(data OrderForm) (OrderForm, error) {
	sql := `UPDATE "orders" SET
	"userId"=COALESCE(NULLIF(:userId, 0),"userId"),
	"orderNumber"=COALESCE(NULLIF(:orderNumber, ''),"orderNumber"),
	"promoId"=COALESCE(NULLIF(:promoId, 0),"promoId"),
	"total"=COALESCE(NULLIF(:total, 0),"total"),
	"deliveryAddress"=COALESCE(NULLIF(:deliveryAddress, ''),"deliveryAddress"),
	"fullName"=COALESCE(NULLIF(:fullName, ''),"fullName"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"priceCut"=COALESCE(NULLIF(:priceCut, 0),"priceCut"),
	"subtotal"=COALESCE(NULLIF(:subtotal, 0),"subtotal"),
	"status"=COALESCE(NULLIF(:status, ''),"status"),
	"deliveryFee"=COALESCE(NULLIF(:deliveryFee, 0),"deliveryFee"),
	"deliveryShipping"=COALESCE(NULLIF(:deliveryShipping, ''),"deliveryShipping"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := OrderForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteOrders(id int) (OrderForm, error) {
	sql := `DELETE FROM "orders" WHERE id = $1 RETURNING *`
	data := OrderForm{}
	err := db.Get(&data, sql, id)
	return data, err
}

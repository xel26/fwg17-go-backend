package models

import (
	"fmt"

	"github.com/LukaGiorgadze/gonull"
)

type OrderProducts struct {
	Id               int    `db:"id" json:"id"`
	Quantity         int    `db:"quantity" json:"quantity"`
	OrderId          int    `db:"orderId" json:"orderId"`
	ProductName      string `db:"productName" json:"productName"`
	Image            string `db:"image" json:"image"`
	BasePrice        int    `db:"basePrice" json:"basePrice"`
	Discount         int    `db:"discount" json:"discount"`
	Tag              gonull.Nullable[string] `db:"tag" json:"tag"`
	Size             string `db:"size" json:"size"`
	Variant          string `db:"variant" json:"variant"`
	DeliveryShipping string `db:"deliveryShipping" json:"deliveryShipping"`
}

type InfoOP struct {
	Data  []OrderProducts
	Count int
}

func CountSubtotal(orderDetailsId int) (OrderDetails, error) {
	sql := `
	update "orderDetails" set "subtotal" = (
        select (("p"."basePrice" - "p"."discount" + "s"."additionalPrice" + "v"."additionalPrice") * "od"."quantity")
        from "orderDetails" "od"
        join "products" "p" on ("p"."id" = "od"."productId")
        join "sizes" "s" on ("s"."id" = "od"."sizeId")
        join "variant" "v" on ("v"."id" = "od"."variantId")
        where "od"."id" = $1
    )
    where "id" = $1
    RETURNING *
	`
	result := OrderDetails{}
	err := db.Get(&result, sql, orderDetailsId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func CountTotal(orderId int) (OrderForm, error) {
	sql := `
	update "orders" set "total" = (select sum("subtotal") from "orderDetails" where "orderId" = $1)
    where "id" = $1
    RETURNING *
	`
	result := OrderForm{}
	err := db.Get(&result, sql, orderId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func CountTax(orderId int) (OrderForm, error) {
	sql := `
    update "orders" set "tax" = (select ("total" * 0.025) from "orders" where "id" = $1)
    where "id" = $1
    RETURNING *
	`
	result := OrderForm{}
	err := db.Get(&result, sql, orderId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func CountTotalTransaction(orderId int) (OrderForm, error) {
	sql := `
    UPDATE "orders" set "subtotal" = (select "total" from "orders" where "id" = $1) + (select "tax" from "orders" where "id" = $1) + (select "deliveryFee" from "orders" where "id" = $1)
    WHERE "id" = $1
    RETURNING *
	`
	result := OrderForm{}
	err := db.Get(&result, sql, orderId)
	if err != nil {
		return result, err
	}

	return result, err
}

func GetAddress(userId int) (UserForm, error) {
	sql := `select "address" from "users" where "id" = $1`
	data := UserForm{}
	err := db.Get(&data, sql, userId)
	return data, err
}

func GetFullName(userId int) (UserForm, error) {
	sql := `select "fullName" from "users" where "id" = $1`
	data := UserForm{}
	err := db.Get(&data, sql, userId)
	return data, err
}

func GetEmail(userId int) (UserForm, error) {
	sql := `select "email" from "users" where "id" = $1`
	data := UserForm{}
	err := db.Get(&data, sql, userId)
	return data, err
}

func FindAllOrdersByUserId(status string, userId int, sortBy string, order string, limit int, offset int) (InfoO, error) {
	fmt.Println(userId)
	sql := `
	SELECT "o".*,
    array_agg(DISTINCT "p"."image") "productsImage"
    FROM "orders" "o"
    JOIN "orderDetails" "od" ON ("od"."orderId" = "o"."id")
    JOIN "products" "p" ON ("p"."id" = "od"."productId")
	WHERE "status" ILIKE $1 AND "userId" = $2
	GROUP BY "o"."id"
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $3 OFFSET $4
	`
	sqlCount := `
	SELECT COUNT(*) FROM "orders"
	WHERE "status" ILIKE $1 AND "userId" = $2
	`

	result := InfoO{}
	data := []Order{}
	error := db.Select(&data, sql, "%"+status+"%", userId, limit, offset)
	if error != nil {
		return result, error
	}

	result.Data = data

	row := db.QueryRow(sqlCount, "%"+status+"%", userId)
	err := row.Scan(&result.Count)

	return result, err
}

func GetOrderProducts(orderId int, userId int, sortBy string, order string) (InfoOP, error) {

	sql := `
	SELECT 
	"od"."id",
	"od"."quantity",
	"od"."orderId",
	"p"."name" AS "productName",
	"p"."image",
	"p"."basePrice",
	"p"."discount",
	"t"."name" AS "tag",
	"s"."size",
	"v"."name" AS "variant",
	"o"."deliveryShipping"
	FROM "orderDetails" "od"
	JOIN "products" "p" on ("p"."id" = "od"."productId")
	LEFT JOIN "tags" "t" on ("t"."id" = "p"."tagId")
	JOIN "sizes" "s" on ("s"."id" = "od"."sizeId")
	JOIN "variant" "v" on ("v"."id" = "od"."variantId")
	JOIN "orders" "o" on ("o"."id" = "od"."orderId")
	WHERE "od"."orderId" = $1
	ORDER BY "`+sortBy+`" `+order+`
	`

	sqlCount := `
	SELECT COUNT(*) FROM "orderDetails" "od"
	JOIN "orders" "o" on ("o"."id" = "od"."orderId")
	WHERE "od"."orderId" = $1
	`

	result := InfoOP{}
	data := []OrderProducts{}
	err := db.Select(&data, sql, orderId)

	if err != nil {
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, orderId)
	err = row.Scan(&result.Count)
	if err != nil {
		return result, err
	}

	return result, err
}

func GetOneSize(size string) (Sizes, error) {
	sql := `SELECT * FROM "sizes" WHERE "size" ILIKE $1`
	data := Sizes{}
	err := db.Get(&data, sql, size)
	return data, err
}

func GetOneVariant(name string) (Variants, error) {
	sql := `SELECT * FROM "variant" WHERE "name" ILIKE $1`
	data := Variants{}
	err := db.Get(&data, sql, name)
	return data, err
}

type OD struct {
	OrderId   int `db:"orderId" json:"orderId" form:"orderId"`
	ProductId int `db:"productId" json:"productId" form:"productId"`
	SizeId    int `db:"sizeId" json:"sizeId" form:"sizeId"`
	VariantId int `db:"variantId" json:"variantId" form:"variantId"`
	Quantity  int `db:"quantity" json:"quantity" form:"quantity"`
}

func CreateOD(data OD) (OrderDetails, error) {
	sql := `
	INSERT INTO "orderDetails" ("orderId", "productId", "sizeId", "variantId", "quantity") 
	VALUES
	(:orderId, :productId, :sizeId, :variantId, :quantity)
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
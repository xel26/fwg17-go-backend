package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id            int            `db:"id" json:"id"`
	Name          string         `db:"name" json:"name"`
	Description   sql.NullString `db:"description" json:"description"`
	Image         sql.NullString `db:"image" json:"image"`
	Discount      sql.NullInt64  `db:"discount" json:"discount"`
	BasePrice     int            `db:"basePrice" json:"basePrice"`
	IsRecommended sql.NullBool   `db:"isRecommended" json:"isRecommended"`
	TagId         sql.NullInt64  `db:"tagId" json:"tagId"`
	CreatedAt     time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type ProductForm struct {
	Id            int          `db:"id" json:"id"`
	Name          *string      `db:"name" json:"name" form:"name" binding:"required,min=3"`
	Description   *string      `db:"description" json:"description" form:"description"`
	Image         string      `db:"image" json:"image"`
	Discount      *int      `db:"discount" json:"discount" form:"discount"`
	BasePrice     *int         `db:"basePrice" json:"basePrice" form:"basePrice" binding:"required,numeric"`
	IsRecommended *bool        `db:"isRecommended" json:"isRecommended" form:"isRecommended"`
	TagId         *int         `db:"tagId" json:"tagId" form:"tagId"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoP struct {
	Data  []Product
	Count int
}

func FindAllProducts(searchKey string, sortBy string, order string, limit int, offset int) (InfoP, error) {
	sql := `
	SELECT * FROM "products" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "products"
	WHERE "name" ILIKE $1
	`

	result := InfoP{}
	data := []Product{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProducts(id int) (Product, error) {
	sql := `SELECT * FROM "products" WHERE id = $1`
	data := Product{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProducts(data ProductForm) (ProductForm, error) {
	sql := `INSERT INTO "products" ("name", "description", "basePrice", "image", "discount", "isRecommended", "tagId") 
	VALUES
	(:name, :description, :basePrice, :image, :discount, :isRecommended, :tagId)
	RETURNING *
	`
	result := ProductForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProduct(data ProductForm) (ProductForm, error) {
	sql := `UPDATE "products" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"description"=COALESCE(NULLIF(:description, ''),"description"),
	"basePrice"=COALESCE(NULLIF(:basePrice, 0),"basePrice"),
	"image"=COALESCE(NULLIF(:image, ''),"image"),
	"discount"=COALESCE(NULLIF(:discount, 0),"discount"),
	"isRecommended"=COALESCE(NULLIF(:isRecommended, false),"isRecommended"),
	"tagId"=COALESCE(NULLIF(:tagId, 0),"tagId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := ProductForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProduct(id int) (ProductForm, error) {
	sql := `DELETE FROM "products" WHERE id = $1 RETURNING *`
	data := ProductForm{}
	err := db.Get(&data, sql, id)
	return data, err
}
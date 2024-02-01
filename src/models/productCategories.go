package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbPC *sqlx.DB = lib.DB

type ProductCategories struct {
	Id         int          `dbPC:"id" json:"id"`
	ProductCategoriesId  int          `dbPC:"productId" json:"productId" form:"productId"`
	CategoryId int          `dbPC:"categoryId" json:"categoryId" form:"categoryId"`
	CreatedAt  time.Time    `dbPC:"createdAt" json:"createdAt"`
	UpdatedAt  sql.NullTime `dbPC:"updatedAt" json:"updatedAt"`
}

type InfoPC struct {
	Data  []ProductCategories
	Count int
}

func FindAllProductCategories(sortBy string, order string, limit int, offset int) (InfoPC, error) {
	sql := `
	SELECT * FROM "productCategories" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "productCategories"
	`

	result := InfoPC{}
	data := []ProductCategories{}
	err := dbPC.Select(&data, sql, limit, offset)
	result.Data = data

	row := dbPC.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductCategories(id int) (ProductCategories, error) {
	sql := `SELECT * FROM "productCategories" WHERE id = $1`
	data := ProductCategories{}
	err := dbPC.Get(&data, sql, id)
	return data, err
}

func CreateProductCategories(data ProductCategories) (ProductCategories, error) {
	sql := `INSERT INTO "productCategories" ("productId", "categoryId") 
	VALUES
	(:productId, :categoryId)
	RETURNING *
	`
	result := ProductCategories{}
	rows, err := dbPC.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductCategories(data ProductCategories) (ProductCategories, error) {
	sql := `UPDATE "productCategories" SET
	"productId"=COALESCE(NULLIF(:productId, ''),"productId"),
	"categoryId"=COALESCE(NULLIF(:categoryId, ''),"categoryId"),
	WHERE id=:id
	RETURNING *
	`
	result := ProductCategories{}
	rows, err := dbPC.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductCategories(id int) (ProductCategories, error) {
	sql := `DELETE FROM "productCategories" WHERE id = $1 RETURNING *`
	data := ProductCategories{}
	err := dbPC.Get(&data, sql, id)
	return data, err
}

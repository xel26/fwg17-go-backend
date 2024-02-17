package models

import (
	"database/sql"
	"time"
)

type ProductCategories struct {
	Id         int          `db:"id" json:"id"`
	ProductId  int          `db:"productId" json:"productId" form:"productId" binding:"required,numeric"`
	CategoryId int          `db:"categoryId" json:"categoryId" form:"categoryId" binding:"required,numeric"`
	CreatedAt  time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt  sql.NullTime `db:"updatedAt" json:"updatedAt"`
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
	err := db.Select(&data, sql, limit, offset)
	if err != nil{
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductCategories(id int) (ProductCategories, error) {
	sql := `SELECT * FROM "productCategories" WHERE id = $1`
	data := ProductCategories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductCategories(data ProductCategories) (ProductCategories, error) {
	sql := `INSERT INTO "productCategories" ("productId", "categoryId") 
	VALUES
	(:productId, :categoryId)
	RETURNING *
	`
	result := ProductCategories{}
	rows, err := db.NamedQuery(sql, data)
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
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"categoryId"=COALESCE(NULLIF(:categoryId, 0),"categoryId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := ProductCategories{}
	rows, err := db.NamedQuery(sql, data)
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
	err := db.Get(&data, sql, id)
	return data, err
}

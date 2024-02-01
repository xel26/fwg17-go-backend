package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbPV *sqlx.DB = lib.DB

type ProductVariants struct {
	Id        int          `dbPV:"id" json:"id"`
	ProductId int          `dbPV:"productId" json:"productId" form:"productId"`
	VariantId int          `dbPV:"variantId" json:"variantId" form:"variantId"`
	CreatedAt time.Time    `dbPV:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `dbPV:"updatedAt" json:"updatedAt"`
}

type InfoPV struct {
	Data  []ProductVariants
	Count int
}

func FindAllProductVariants(sortBy string, order string, limit int, offset int) (InfoPV, error) {
	sql := `
	SELECT * FROM "productVariant" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "productVariant"
	`

	result := InfoPV{}
	data := []ProductVariants{}
	err := dbPV.Select(&data, sql, limit, offset)
	result.Data = data

	row := dbPV.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductVariants(id int) (ProductVariants, error) {
	sql := `SELECT * FROM "productVariant" WHERE id = $1`
	data := ProductVariants{}
	err := dbPV.Get(&data, sql, id)
	return data, err
}

func CreateProductVariants(data ProductVariants) (ProductVariants, error) {
	sql := `INSERT INTO "productVariant" ("productId", "variantId") 
	VALUES
	(:productId, :variantId)
	RETURNING *
	`
	result := ProductVariants{}
	rows, err := dbPV.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductVariants(data ProductVariants) (ProductVariants, error) {
	sql := `UPDATE "productVariant" SET
	"productId"=COALESCE(NULLIF(:productId, ''),"productId"),
	"variantId"=COALESCE(NULLIF(:variantId, ''),"variantId"),
	WHERE id=:id
	RETURNING *
	`
	result := ProductVariants{}
	rows, err := dbPV.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductVariants(id int) (ProductVariants, error) {
	sql := `DELETE FROM "productVariant" WHERE id = $1 RETURNING *`
	data := ProductVariants{}
	err := dbPV.Get(&data, sql, id)
	return data, err
}
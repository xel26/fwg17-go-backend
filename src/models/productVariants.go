package models

import (
	"database/sql"
	"time"
)

type ProductVariants struct {
	Id        int          `db:"id" json:"id"`
	ProductId int          `db:"productId" json:"productId" form:"productId"`
	VariantId int          `db:"variantId" json:"variantId" form:"variantId"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
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
	err := db.Select(&data, sql, limit, offset)
	
	if err != nil{
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductVariants(id int) (ProductVariants, error) {
	sql := `SELECT * FROM "productVariant" WHERE id = $1`
	data := ProductVariants{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductVariants(data ProductVariants) (ProductVariants, error) {
	sql := `INSERT INTO "productVariant" ("productId", "variantId") 
	VALUES
	(:productId, :variantId)
	RETURNING *
	`
	result := ProductVariants{}
	rows, err := db.NamedQuery(sql, data)
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
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"variantId"=COALESCE(NULLIF(:variantId, 0),"variantId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := ProductVariants{}
	rows, err := db.NamedQuery(sql, data)
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
	err := db.Get(&data, sql, id)
	return data, err
}
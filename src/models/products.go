package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbP *sqlx.DB = lib.DB

type Product struct {
	Id            int            `dbP:"id" json:"id"`
	Name          string         `dbP:"name" json:"name" form:"name"`
	Description   sql.NullString `dbP:"description" json:"description" form:"description"`
	Image         sql.NullString `dbP:"image" json:"image" form:"image"`
	Discount      sql.NullInt64  `dbP:"discount" json:"discount" form:"discount"`
	BasePrice     int            `dbP:"basePrice" json:"basePrice" form:"basePrice"`
	IsRecommended sql.NullBool   `dbP:"isRecommended" json:"isRecommended" form:"isRecommended"`
	TagId         sql.NullInt64  `dbP:"tagId" json:"tagId" form:"tagId"`
	CreatedAt     time.Time      `dbP:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `dbP:"updatedAt" json:"updatedAt"`
}


type InfoP struct {
	Data  []Product
	Count int
}




func FindAllProducts(searchKey string, sortBy string, order string, limit int, offset int) (InfoP, error) {
	sql := `
	SELECT * FROM "products" 
	WHERE "name" ILIKE $1
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "products"
	WHERE "name" ILIKE $1
	`

	result := InfoP{}
	data := []Product{}
	err := dbP.Select(&data, sql,"%"+searchKey+"%", limit, offset)
	result.Data = data
	
	row := dbP.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}



func FindOneProducts(id int) (Product, error) {
	sql := `SELECT * FROM "products" WHERE id = $1`
	data := Product{}
	err := dbP.Get(&data, sql, id)
	return data, err
}




func CreateProducts(data Product) (Product, error) {
	sql := `INSERT INTO "products" ("name", "description", "basePrice", "image", "discount", "isRecommended", "tagId") 
	VALUES
	(:name, :description, :basePrice, :image, :discount, :isRecommended, :tagId)
	RETURNING *
	`
	result := Product{}
	rows, err := dbP.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}
	
	for rows.Next(){
		rows.StructScan(&result)
	}
	
	return result, err
}



func UpdateProduct(data Product) (Product, error) {
	sql := `UPDATE "products" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"description"=COALESCE(NULLIF(:description, ''),"description"),
	"basePrice"=COALESCE(NULLIF(:basePrice, ''),"basePrice"),
	"image"=COALESCE(NULLIF(:image, ''),"image"),
	"discount"=COALESCE(NULLIF(:discount, ''),"discount"),
	"isRecommended"=COALESCE(NULLIF(:isRecommended, ''),"isRecommended"),
	"tagId"=COALESCE(NULLIF(:tagId, ''),"tagId")
	WHERE id=:id
	RETURNING *
	`
	result := Product{}
	rows, err := dbP.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next(){
		rows.StructScan(&result)
	}

	return result, err
}



func DeleteProduct(id int) (Product, error) {
	sql := `DELETE FROM "products" WHERE id = $1 RETURNING *`
	data := Product{}
	err := dbP.Get(&data, sql, id)
	return data, err
}
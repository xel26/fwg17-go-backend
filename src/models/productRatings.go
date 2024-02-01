package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbPR *sqlx.DB = lib.DB

type ProductRatings struct {
	Id            int            `dbPR:"id" json:"id"`
	ProductId     int            `dbPR:"productId" json:"productId" form:"productId"`
	Rate          int            `dbPR:"rate" json:"rate" form:"rate"`
	ReviewMessage sql.NullString `dbPR:"reviewMessage" json:"reviewMessage" form:"reviewMessage"`
	UserId        int            `dbPR:"userId" json:"userId" form:"userId"`
	CreatedAt     time.Time      `dbPR:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `dbPR:"updatedAt" json:"updatedAt"`
}

type InfoPR struct {
	Data  []ProductRatings
	Count int
}

func FindAllProductRatings(sortBy string, order string, limit int, offset int) (InfoPR, error) {
	sql := `
	SELECT * FROM "productRatings" 
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "productRatings"
	`

	result := InfoPR{}
	data := []ProductRatings{}
	err := dbPR.Select(&data, sql, limit, offset)
	result.Data = data

	row := dbPR.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductRatings(id int) (ProductRatings, error) {
	sql := `SELECT * FROM "productRatings" WHERE id = $1`
	data := ProductRatings{}
	err := dbPR.Get(&data, sql, id)
	return data, err
}

func CreateProductRatings(data ProductRatings) (ProductRatings, error) {
	sql := `INSERT INTO "productRatings" ("productId", "rate", "reviewMessage", "userId") 
	VALUES
	(:productId, :rate, :reviewMessage, :userId)
	RETURNING *
	`
	result := ProductRatings{}
	rows, err := dbPR.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductRatings(data ProductRatings) (ProductRatings, error) {
	sql := `UPDATE "productRatings" SET
	"productId"=COALESCE(NULLIF(:productId, ''),"productId"),
	"rate"=COALESCE(NULLIF(:rate, ''),"rate"),
	"reviewMessage"=COALESCE(NULLIF(:reviewMessage, ''),"reviewMessage"),
	"userId"=COALESCE(NULLIF(:userId, ''),"userId"),
	WHERE id=:id
	RETURNING *
	`
	result := ProductRatings{}
	rows, err := dbPR.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductRatings(id int) (ProductRatings, error) {
	sql := `DELETE FROM "productRatings" WHERE id = $1 RETURNING *`
	data := ProductRatings{}
	err := dbPR.Get(&data, sql, id)
	return data, err
}

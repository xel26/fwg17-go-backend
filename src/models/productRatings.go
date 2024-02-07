package models

import (
	"database/sql"
	"time"
)

type ProductRatings struct {
	Id            int            `db:"id" json:"id"`
	ProductId     int            `db:"productId" json:"productId"`
	Rate          int            `db:"rate" json:"rate"`
	ReviewMessage sql.NullString `db:"reviewMessage" json:"reviewMessage"`
	UserId        int            `db:"userId" json:"userId"`
	CreatedAt     time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type PRForm struct {
	Id            int          `db:"id" json:"id"`
	ProductId     *int         `db:"productId" json:"productId" form:"productId" binding:"required,numeric"`
	Rate          *int         `db:"rate" json:"rate" form:"rate" binding:"required,eq=5|eq=4|eq=3|eq=2|eq=1"`
	ReviewMessage *string      `db:"reviewMessage" json:"reviewMessage" form:"reviewMessage"`
	UserId        *int         `db:"userId" json:"userId" form:"userId" binding:"required,numeric"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
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
	err := db.Select(&data, sql, limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneProductRatings(id int) (ProductRatings, error) {
	sql := `SELECT * FROM "productRatings" WHERE id = $1`
	data := ProductRatings{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateProductRatings(data PRForm) (PRForm, error) {
	sql := `INSERT INTO "productRatings" ("productId", "rate", "reviewMessage", "userId") 
	VALUES
	(:productId, :rate, :reviewMessage, :userId)
	RETURNING *
	`
	result := PRForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateProductRatings(data PRForm) (PRForm, error) {
	sql := `UPDATE "productRatings" SET
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"rate"=COALESCE(NULLIF(:rate, 0),"rate"),
	"reviewMessage"=COALESCE(NULLIF(:reviewMessage, ''),"reviewMessage"),
	"userId"=COALESCE(NULLIF(:userId, 0),"userId"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := PRForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteProductRatings(id int) (PRForm, error) {
	sql := `DELETE FROM "productRatings" WHERE id = $1 RETURNING *`
	data := PRForm{}
	err := db.Get(&data, sql, id)
	return data, err
}

package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbS *sqlx.DB = lib.DB

type Sizes struct {
	Id              int           `dbS:"id" json:"id"`
	Size            string        `dbS:"size" json:"size" form:"size"`
	AdditionalPrice sql.NullInt64 `dbS:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       time.Time     `dbS:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime  `dbS:"updatedAt" json:"updatedAt"`
}

type InfoS struct {
	Data  []Sizes
	Count int
}

func FindAllSizes(searchKey string, sortBy string, order string, limit int, offset int) (InfoS, error) {
	sql := `
	SELECT * FROM "sizes" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "sizes"
	WHERE "name" ILIKE $1
	`

	result := InfoS{}
	data := []Sizes{}
	err := dbS.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := dbS.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneSizes(id int) (Sizes, error) {
	sql := `SELECT * FROM "sizes" WHERE id = $1`
	data := Sizes{}
	err := dbS.Get(&data, sql, id)
	return data, err
}

func CreateSizes(data Sizes) (Sizes, error) {
	sql := `INSERT INTO "sizes" ("size", "additionalPrice") 
	VALUES
	(:size, :additionalPrice)
	RETURNING *
	`
	result := Sizes{}
	rows, err := dbS.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateSizes(data Sizes) (Sizes, error) {
	sql := `UPDATE "sizes" SET
	"size"=COALESCE(NULLIF(:size, ''),"size"),
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice, ''),"additionalPrice")
	WHERE id=:id
	RETURNING *
	`
	result := Sizes{}
	rows, err := dbS.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteSizes(id int) (Sizes, error) {
	sql := `DELETE FROM "sizes" WHERE id = $1 RETURNING *`
	data := Sizes{}
	err := dbS.Get(&data, sql, id)
	return data, err
}

package models

import (
	"database/sql"
	"time"
)

type Sizes struct {
	Id              int           `db:"id" json:"id"`
	Size            string        `db:"size" json:"size" form:"size"`
	AdditionalPrice sql.NullInt64 `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       time.Time     `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime  `db:"updatedAt" json:"updatedAt"`
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
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneSizes(id int) (Sizes, error) {
	sql := `SELECT * FROM "sizes" WHERE id = $1`
	data := Sizes{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateSizes(data Sizes) (Sizes, error) {
	sql := `INSERT INTO "sizes" ("size", "additionalPrice") 
	VALUES
	(:size, :additionalPrice)
	RETURNING *
	`
	result := Sizes{}
	rows, err := db.NamedQuery(sql, data)
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
	rows, err := db.NamedQuery(sql, data)
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
	err := db.Get(&data, sql, id)
	return data, err
}

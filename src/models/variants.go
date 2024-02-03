package models

import (
	"database/sql"
	"time"
)

type Variants struct {
	Id              int          `db:"id" json:"id"`
	Name            string       `db:"name" json:"name" form:"name"`
	AdditionalPrice int          `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoV struct {
	Data  []Variants
	Count int
}

func FindAllVariants(searchKey string, sortBy string, order string, limit int, offset int) (InfoV, error) {
	sql := `
	SELECT * FROM "variant" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "variant"
	WHERE "name" ILIKE $1
	`

	result := InfoV{}
	data := []Variants{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneVariants(id int) (Variants, error) {
	sql := `SELECT * FROM "variant" WHERE id = $1`
	data := Variants{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateVariants(data Variants) (Variants, error) {
	sql := `INSERT INTO "variant" ("name", "additionalPrice") VALUES (:name, :additionalPrice) RETURNING *`
	result := Variants{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateVariants(data Variants) (Variants, error) {
	sql := `UPDATE "variant" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"updatedAt" NOW()
	WHERE id=:id
	RETURNING *
	`
	result := Variants{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteVariants(id int) (Variants, error) {
	sql := `DELETE FROM "variant" WHERE id = $1 RETURNING *`
	data := Variants{}
	err := db.Get(&data, sql, id)
	return data, err
}

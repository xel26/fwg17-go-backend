package models

import (
	"database/sql"
	"time"
)

type Sizes struct {
	Id              int           `db:"id" json:"id"`
	Size            string        `db:"size" json:"size" form:"size"`
	AdditionalPrice int `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice"`
	CreatedAt       time.Time     `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime  `db:"updatedAt" json:"updatedAt"`
}

type SizesForm struct {
	Id              int           `db:"id" json:"id"`
	Size            string        `db:"size" json:"size" form:"size" binding:"required,min=3"`
	AdditionalPrice *int `db:"additionalPrice" json:"additionalPrice" form:"additionalPrice" binding:"required"`
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
	WHERE "size" ILIKE $1
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "sizes"
	WHERE "size" ILIKE $1
	`

	result := InfoS{}
	data := []Sizes{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil{
		return result, err
	}
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

func CreateSizes(data SizesForm) (SizesForm, error) {
	sql := `INSERT INTO "sizes" ("size", "additionalPrice") 
	VALUES
	(:size, :additionalPrice)
	RETURNING *
	`
	result := SizesForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateSizes(data SizesForm) (SizesForm, error) {
	sql := `UPDATE "sizes" SET
	"size"=COALESCE(NULLIF(:size, ''),"size"),
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice, 0),"additionalPrice"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := SizesForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteSizes(id int) (SizesForm, error) {
	sql := `DELETE FROM "sizes" WHERE id = $1 RETURNING *`
	data := SizesForm{}
	err := db.Get(&data, sql, id)
	return data, err
}

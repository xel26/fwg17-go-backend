package models

import (
	"database/sql"
	"time"
)

type Promo struct {
	Id            int            `db:"id" json:"id"`
	Name          string         `db:"name" json:"name" form:"name"`
	Code          string         `db:"code" json:"code" form:"code"`
	Description   sql.NullString `db:"description" json:"description" form:"description"`
	Percentage    float64        `db:"percentage" json:"percentage" form:"percentage"`
	IsExpired     sql.NullBool   `db:"isExpired" json:"isExpired" form:"isExpired"`
	MaximumPromo  int            `db:"maximumPromo" json:"maximumPromo" form:"maximumPromo"`
	MinimumAmount int            `db:"minimumAmount" json:"minimumAmount" form:"minimumAmount"`
	CreatedAt     time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type PromoForm struct {
	Id            int          `db:"id" json:"id"`
	Name          *string      `db:"name" json:"name" form:"name"`
	Code          *string      `db:"code" json:"code" form:"code"`
	Description   *string      `db:"description" json:"description" form:"description"`
	Percentage    *float64     `db:"percentage" json:"percentage" form:"percentage"`
	IsExpired     *bool        `db:"isExpired" json:"isExpired" form:"isExpired"`
	MaximumPromo  *int         `db:"maximumPromo" json:"maximumPromo" form:"maximumPromo"`
	MinimumAmount *int         `db:"minimumAmount" json:"minimumAmount" form:"minimumAmount"`
	CreatedAt     time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoPo struct {
	Data  []Promo
	Count int
}

func FindAllPromo(searchKey string, sortBy string, order string, limit int, offset int) (InfoPo, error) {
	sql := `
	SELECT * FROM "promo" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "promo"
	WHERE "name" ILIKE $1
	`

	result := InfoPo{}
	data := []Promo{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOnePromo(id int) (Promo, error) {
	sql := `SELECT * FROM "promo" WHERE id = $1`
	data := Promo{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreatePromo(data PromoForm) (PromoForm, error) {
	sql := `INSERT INTO "promo" ("name","code", "description", "percentage", "isExpired", "maximumPromo", "minimumAmount") 
	VALUES
	(:name, :code, :description, :percentage, :isExpired, :maximumPromo, :minimumAmount)
	RETURNING *
	`
	result := PromoForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdatePromo(data PromoForm) (PromoForm, error) {
	sql := `UPDATE "promo" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"code"=COALESCE(NULLIF(:code, ''),"code"),
	"description"=COALESCE(NULLIF(:description, ''),"description"),
	"percentage"=COALESCE(NULLIF(:percentage, 0),"percentage"),
	"isExpired"=COALESCE(NULLIF(:isExpired, false),"isExpired"),
	"maximumPromo"=COALESCE(NULLIF(:maximumPromo, 0),"maximumPromo"),
	"minimumAmount"=COALESCE(NULLIF(:minimumAmount, 0),"minimumAmount"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := PromoForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeletePromo(id int) (PromoForm, error) {
	sql := `DELETE FROM "promo" WHERE id = $1 RETURNING *`
	data := PromoForm{}
	err := db.Get(&data, sql, id)
	return data, err
}

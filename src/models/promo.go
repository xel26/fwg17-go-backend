package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbPo *sqlx.DB = lib.DB

type Promo struct {
	Id            int            `dbPo:"id" json:"id"`
	Name          string         `dbPo:"name" json:"name" form:"name"`
	Code          string         `dbPo:"code" json:"code" form:"code"`
	Description   sql.NullString `dbPo:"description" json:"description" form:"description"`
	Percentage    float64        `dbPo:"percentage" json:"percentage" form:"percentage"`
	IsExpired     sql.NullBool   `dbPo:"isExpired" json:"isExpired" form:"isExpired"`
	MaximumPromo  int            `dbPo:"maximumPromo" json:"maximumPromo" form:"maximumPromo"`
	MinimumAmount int            `dbPo:"minimumAmount" json:"minimumAmount" form:"minimumAmount"`
	CreatedAt     time.Time      `dbPo:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `dbPo:"updatedAt" json:"updatedAt"`
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
	err := dbPo.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := dbPo.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOnePromo(id int) (Promo, error) {
	sql := `SELECT * FROM "promo" WHERE id = $1`
	data := Promo{}
	err := dbPo.Get(&data, sql, id)
	return data, err
}

func CreatePromo(data Promo) (Promo, error) {
	sql := `INSERT INTO "promo" ("name","code", "description", "percentage", "isExpired", "maximumPromo", "minimumAmount") 
	VALUES
	(:name, :code, :description, :percentage, :isExpired, :maximumPromo, :minimumAmount)
	RETURNING *
	`
	result := Promo{}
	rows, err := dbPo.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdatePromo(data Promo) (Promo, error) {
	sql := `UPDATE "promo" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"code"=COALESCE(NULLIF(:code, ''),"code"),
	"description"=COALESCE(NULLIF(:description, ''),"description"),
	"percentage"=COALESCE(NULLIF(:percentage, ''),"percentage"),
	"isExpired"=COALESCE(NULLIF(:isExpired, ''),"isExpired"),
	"maximumPromo"=COALESCE(NULLIF(:maximumPromo, ''),"maximumPromo"),
	"minimumAmount"=COALESCE(NULLIF(:minimumAmount, ''),"minimumAmount")
	WHERE id=:id
	RETURNING *
	`
	result := Promo{}
	rows, err := dbPo.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeletePromo(id int) (Promo, error) {
	sql := `DELETE FROM "promo" WHERE id = $1 RETURNING *`
	data := Promo{}
	err := dbPo.Get(&data, sql, id)
	return data, err
}

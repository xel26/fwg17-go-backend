package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbV *sqlx.DB = lib.DB

type Variants struct {
	Id            int            `dbV:"id" json:"id"`
	Name          string         `dbV:"name" json:"name" form:"name"`
	CreatedAt     time.Time      `dbV:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `dbV:"updatedAt" json:"updatedAt"`
}


type InfoV struct {
	Data  []Variants
	Count int
}


func FindAllVariants(searchKey string, sortBy string, order string, limit int, offset int) (InfoV, error) {
	sql := `
	SELECT * FROM "variant" 
	WHERE "name" ILIKE $1
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "variant"
	WHERE "name" ILIKE $1
	`

	result := InfoV{}
	data := []Variants{}
	err := dbV.Select(&data, sql,"%"+searchKey+"%", limit, offset)
	result.Data = data
	
	row := dbV.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}


func FindOneVariants(id int) (Variants, error) {
	sql := `SELECT * FROM "variant" WHERE id = $1`
	data := Variants{}
	err := dbV.Get(&data, sql, id)
	return data, err
}


func CreateVariants(data Variants) (Variants, error) {
	sql := `INSERT INTO "variant" ("name") VALUES (:name) RETURNING *`
	result := Variants{}
	rows, err := dbV.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}
	
	for rows.Next(){
		rows.StructScan(&result)
	}
	
	return result, err
}



func UpdateVariants(data Variants) (Variants, error) {
	sql := `UPDATE "variant" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	WHERE id=:id
	RETURNING *
	`
	result := Variants{}
	rows, err := dbV.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next(){
		rows.StructScan(&result)
	}

	return result, err
}



func DeleteVariants(id int) (Variants, error) {
	sql := `DELETE FROM "variant" WHERE id = $1 RETURNING *`
	data := Variants{}
	err := dbV.Get(&data, sql, id)
	return data, err
}
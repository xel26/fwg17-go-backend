package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbC *sqlx.DB = lib.DB

type Categories struct {
	Id            int            `dbC:"id" json:"id"`
	Name          string         `dbC:"name" json:"name" form:"name"`
	CreatedAt     time.Time      `dbC:"createdAt" json:"createdAt"`
	UpdatedAt     sql.NullTime   `dbC:"updatedAt" json:"updatedAt"`
}


type InfoC struct {
	Data  []Categories
	Count int
}


func FindAllCategories(searchKey string, sortBy string, order string, limit int, offset int) (InfoC, error) {
	sql := `
	SELECT * FROM "categories" 
	WHERE "name" ILIKE $1
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "categories"
	WHERE "name" ILIKE $1
	`

	result := InfoC{}
	data := []Categories{}
	err := dbC.Select(&data, sql,"%"+searchKey+"%", limit, offset)
	result.Data = data
	
	row := dbC.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}


func FindOneCategories(id int) (Categories, error) {
	sql := `SELECT * FROM "categories" WHERE id = $1`
	data := Categories{}
	err := dbC.Get(&data, sql, id)
	return data, err
}


func CreateCategories(data Categories) (Categories, error) {
	sql := `INSERT INTO "categories" ("name") VALUES (:name) RETURNING *`
	result := Categories{}
	rows, err := dbC.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}
	
	for rows.Next(){
		rows.StructScan(&result)
	}
	
	return result, err
}



func UpdateCategories(data Categories) (Categories, error) {
	sql := `UPDATE "categories" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	WHERE id=:id
	RETURNING *
	`
	result := Categories{}
	rows, err := dbC.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next(){
		rows.StructScan(&result)
	}

	return result, err
}



func DeleteCategories(id int) (Categories, error) {
	sql := `DELETE FROM "categories" WHERE id = $1 RETURNING *`
	data := Categories{}
	err := dbC.Get(&data, sql, id)
	return data, err
}
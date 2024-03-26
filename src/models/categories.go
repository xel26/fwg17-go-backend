package models

import (
	"fmt"
	"time"

	"github.com/LukaGiorgadze/gonull"
)

type Categories struct {
	Id        int          `db:"id" json:"id"`
	Name      *string      `db:"name" json:"name" form:"name"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}

type InfoC struct {
	Data  []Categories
	Count int
}

func FindAllCategories(searchKey string, sortBy string, order string, limit int, offset int) (InfoC, error) {
	sql := `
	SELECT * FROM "categories" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "categories"
	WHERE "name" ILIKE $1
	`

	result := InfoC{}
	data := []Categories{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil{
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneCategories(id int) (Categories, error) {
	sql := `SELECT * FROM "categories" WHERE id = $1`
	data := Categories{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateCategories(data Categories) (Categories, error) {
	sql := `INSERT INTO "categories" ("name") VALUES (:name) RETURNING *`
	result := Categories{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateCategories(data Categories) (Categories, error) {
	sql := `UPDATE "categories" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := Categories{}
	rows, err := db.NamedQuery(sql, data)
	fmt.Println(err)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteCategories(id int) (Categories, error) {
	sql := `DELETE FROM "categories" WHERE id = $1 RETURNING *`
	data := Categories{}
	err := db.Get(&data, sql, id)
	return data, err
}

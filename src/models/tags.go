package models

import (
	"database/sql"
	"time"
)

type Tags struct {
	Id              int           `db:"id" json:"id"`
	Name            string        `db:"name" json:"name" form:"name"`
	CreatedAt       time.Time     `db:"createdAt" json:"createdAt"`
	UpdatedAt       sql.NullTime  `db:"updatedAt" json:"updatedAt"`
}

type InfoT struct {
	Data  []Tags
	Count int
}

func FindAllTags(searchKey string, sortBy string, order string, limit int, offset int) (InfoT, error) {
	sql := `
	SELECT * FROM "tags" 
	WHERE "name" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "tags"
	WHERE "name" ILIKE $1
	`

	result := InfoT{}
	data := []Tags{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneTags(id int) (Tags, error) {
	sql := `SELECT * FROM "tags" WHERE id = $1`
	data := Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateTags(data Tags) (Tags, error) {
	sql := `INSERT INTO "tags" ("name") VALUES (:name)
	RETURNING *
	`
	result := Tags{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateTags(data Tags) (Tags, error) {
	sql := `UPDATE "tags" SET
	"name"=COALESCE(NULLIF(:name, ''),"name")
	WHERE id=:id
	RETURNING *
	`
	result := Tags{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteTags(id int) (Tags, error) {
	sql := `DELETE FROM "tags" WHERE id = $1 RETURNING *`
	data := Tags{}
	err := db.Get(&data, sql, id)
	return data, err
}

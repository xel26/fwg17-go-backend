package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbTs *sqlx.DB = lib.DB

type Testimonial struct {
	Id        int            `dbTs:"id" json:"id"`
	FullName  string         `dbTs:"fullName" json:"fullName" form:"fullName"`
	Role      string         `dbTs:"role" json:"role" form:"role"`
	FeedBack  string         `dbTs:"feedBack" json:"feedBack" form:"feedBack"`
	Image     sql.NullString `dbTs:"image" json:"image" form:"image"`
	Rate      int            `dbTs:"rate" json:"rate" form:"rate"`
	CreatedAt time.Time      `dbTs:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime   `dbTs:"updatedAt" json:"updatedAt"`
}

type InfoTs struct {
	Data  []Testimonial
	Count int
}

func FindAllTestimonial(searchKey string, sortBy string, order string, limit int, offset int) (InfoTs, error) {
	sql := `
	SELECT * FROM "testimonial" 
	WHERE "fullName" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "testimonial"
	WHERE "fullName" ILIKE $1
	`

	result := InfoTs{}
	data := []Testimonial{}
	err := dbTs.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := dbTs.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneTestimonial(id int) (Testimonial, error) {
	sql := `SELECT * FROM "testimonial" WHERE id = $1`
	data := Testimonial{}
	err := dbTs.Get(&data, sql, id)
	return data, err
}

func CreateTestimonial(data Testimonial) (Testimonial, error) {
	sql := `INSERT INTO "testimonial" ("fullName", "role", "feedBack", "rate", "image") 
	VALUES
	(:fullName, :role, :feedBack, :rate, :image)
	RETURNING *
	`
	result := Testimonial{}
	rows, err := dbTs.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateTestimonial(data Testimonial) (Testimonial, error) {
	sql := `UPDATE "testimonial" SET
	"fullName"=COALESCE(NULLIF(:fullName, ''),"fullName"),
	"role"=COALESCE(NULLIF(:role, ''),"role"),
	"feedBack"=COALESCE(NULLIF(:feedBack, ''),"feedBack"),
	"rate"=COALESCE(NULLIF(:rate, ''),"rate"),
	"image"=COALESCE(NULLIF(:image, ''),"image"),
	WHERE id=:id
	RETURNING *
	`
	result := Testimonial{}
	rows, err := dbTs.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteTestimonial(id int) (Testimonial, error) {
	sql := `DELETE FROM "testimonial" WHERE id = $1 RETURNING *`
	data := Testimonial{}
	err := dbTs.Get(&data, sql, id)
	return data, err
}

package models

import (
	"database/sql"
	"time"
)

type Testimonial struct {
	Id        int            `db:"id" json:"id"`
	FullName  string         `db:"fullName" json:"fullName" form:"fullName"`
	Role      string         `db:"role" json:"role" form:"role"`
	Feedback  string         `db:"feedback" json:"feedback" form:"feedback"`
	Image     sql.NullString `db:"image" json:"image" form:"image"`
	Rate      int            `db:"rate" json:"rate" form:"rate"`
	CreatedAt time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type InfoTs struct {
	Data  []Testimonial
	Count int
}

func FindAllTestimonial(searchKey string, sortBy string, order string, limit int, offset int) (InfoTs, error) {
	sql := `
	SELECT * FROM "testimonial" 
	WHERE "fullName" ILIKE $1
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "testimonial"
	WHERE "fullName" ILIKE $1
	`

	result := InfoTs{}
	data := []Testimonial{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneTestimonial(id int) (Testimonial, error) {
	sql := `SELECT * FROM "testimonial" WHERE id = $1`
	data := Testimonial{}
	err := db.Get(&data, sql, id)
	return data, err
}

func CreateTestimonial(data Testimonial) (Testimonial, error) {
	sql := `INSERT INTO "testimonial" ("fullName", "role", "feedback", "rate", "image") 
	VALUES
	(:fullName, :role, :feedback, :rate, :image)
	RETURNING *
	`
	result := Testimonial{}
	rows, err := db.NamedQuery(sql, data)
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
	"feedback"=COALESCE(NULLIF(:feedback, ''),"feedback"),
	"rate"=COALESCE(NULLIF(:rate, ''),"rate"),
	"image"=COALESCE(NULLIF(:image, ''),"image"),
	"updatedAt" NOW()
	WHERE id=:id
	RETURNING *
	`
	result := Testimonial{}
	rows, err := db.NamedQuery(sql, data)
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
	err := db.Get(&data, sql, id)
	return data, err
}

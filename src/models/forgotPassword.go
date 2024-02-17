package models

import (
	"database/sql"
	"time"
)

type ForgotPassword struct {
	Id        int           `db:"id" json:"id"`
	Otp       string        `db:"otp" json:"otp"`
	UserId    sql.NullInt64 `db:"userId" json:"userId"`
	Email     string        `db:"email" json:"email"`
	CreatedAt time.Time     `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime  `db:"updatedAt" json:"updatedAt"`
}
type FPForm struct {
	Id        int          `db:"id" json:"id"`
	Otp       *string      `db:"otp" json:"otp" form:"otp" binding:"required"`
	UserId    *int         `db:"userId" json:"userId" form:"userId" binding:"required"`
	Email     *string      `db:"email" json:"email" form:"email" binding:"required"`
	CreatedAt time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type InfoFP struct {
	Data  []ForgotPassword
	Count int
}

func FindAllForgotPassword(searchKey string, sortBy string, order string, limit int, offset int) (InfoFP, error) {
	sql := `
	SELECT * FROM "forgotPassword" 
	WHERE "email" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "forgotPassword"
	WHERE "email" ILIKE $1
	`

	result := InfoFP{}
	data := []ForgotPassword{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	if err != nil{
		return result, err
	}
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneForgotPassword(id int) (ForgotPassword, error) {
	sql := `SELECT * FROM "forgotPassword" WHERE id = $1`
	data := ForgotPassword{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneByOtp(otp string) (ForgotPassword, error) {
	sql := `SELECT * FROM "forgotPassword" WHERE otp = $1`
	data := ForgotPassword{}
	err := db.Get(&data, sql, otp)
	return data, err
}

func CreateForgotPassword(data ForgotPassword) (ForgotPassword, error) {
	sql := `INSERT INTO "forgotPassword" ("otp", "email", "userId") 
	VALUES (:otp, :email, :userId) 
	RETURNING *
	`
	result := ForgotPassword{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateForgotPassword(data FPForm) (FPForm, error) {
	sql := `UPDATE "forgotPassword" SET
	"otp"=COALESCE(NULLIF(:otp, 0),"otp"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"updatedAt"=CURRENT_TIMESTAMP
	WHERE id=:id
	RETURNING *
	`
	result := FPForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteForgotPassword(id int) (FPForm, error) {
	sql := `DELETE FROM "forgotPassword" WHERE id = $1 RETURNING *`
	data := FPForm{}
	err := db.Get(&data, sql, id)
	return data, err
}

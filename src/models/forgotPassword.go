package models

import (
	"database/sql"
	"time"
)

type ForgotPassword struct {
	Id        int            `db:"id" json:"id"`
	Otp       sql.NullString `db:"otp" json:"otp" form:"otp"`
	Email     sql.NullString `db:"email" json:"email" form:"email"`
	UserId    sql.NullInt64  `db:"userId" json:"userId" form:"userId"`
	CreatedAt time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type InfoFP struct {
	Data  []ForgotPassword
	Count int
}


func FindAllForgotPassword(searchKey string, sortBy string, order string, limit int, offset int) (InfoFP, error) {
	sql := `
	SELECT * FROM "forgotPassword" 
	WHERE "email" ILIKE $1
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "forgotPassword"
	WHERE "email" ILIKE $1
	`

	result := InfoFP{}
	data := []ForgotPassword{}
	err := db.Select(&data, sql,"%"+searchKey+"%", limit, offset)
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


func UpdateForgotPassword(data ForgotPassword) (ForgotPassword, error) {
	sql := `UPDATE "forgotPassword" SET
	"otp"=COALESCE(NULLIF(:otp, ''),"otp"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"userId"=COALESCE(NULLIF(:userId, ''),"userId")
	WHERE id=:id
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

func DeleteForgotPassword(id int) (ForgotPassword, error) {
	sql := `DELETE FROM "forgotPassword" WHERE id = $1 RETURNING *`
	data := ForgotPassword{}
	err := db.Get(&data, sql, id)
	return data, err
}

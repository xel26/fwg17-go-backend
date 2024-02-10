package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id          int            `db:"id" json:"id"`
	FullName    string         `db:"fullName" json:"fullName" form:"fullName"`
	Email       string         `db:"email" json:"email" form:"email" form:"email"`
	Password    string         `db:"password" json:"-" form:"password" form:"password"`
	Address     sql.NullString `db:"address" json:"address" form:"address"`
	Picture     string         `db:"picture" json:"picture"`
	PhoneNumber sql.NullString `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        string         `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time      `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime   `db:"updatedAt" json:"updatedAt"`
}

type UserForm struct {
	Id          int          `db:"id" json:"id"`
	FullName    *string      `db:"fullName" json:"fullName" form:"fullName" binding:"required,min=3"`
	Email       *string      `db:"email" json:"email" form:"email" binding:"email,required"`
	Password    string       `db:"password" json:"-" form:"password" binding:"required"`
	Address     *string      `db:"address" json:"address" form:"address"`
	Picture     string       `db:"picture" json:"picture"`
	PhoneNumber *string      `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        *string      `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type Info struct {
	Data  []User
	Count int
}

func FindAllUsers(searchKey string, sortBy string, order string, limit int, offset int) (Info, error) {
	sql := `
	SELECT * FROM "users" 
	WHERE "fullName" ILIKE $1
	ORDER BY "` + sortBy + `" ` + order + `
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "users"
	WHERE "fullName" ILIKE $1
	`

	result := Info{}
	data := []User{}
	err := db.Select(&data, sql, "%"+searchKey+"%", limit, offset)
	result.Data = data

	row := db.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}

func FindOneUsers(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE id = $1`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}

func FindOneUsersByEmail(email string) (User, error) {
	sql := `SELECT * FROM "users" WHERE email = $1`
	data := User{}
	err := db.Get(&data, sql, email)
	return data, err
}

func CreateUser(data UserForm) (UserForm, error) {
	fmt.Println(data.PhoneNumber)
	fmt.Println(data.Picture)

	sql := `INSERT INTO "users" ("fullName", "email", "password", "address", "phoneNumber", "role", "picture") 
	VALUES
	(:fullName, :email, :password, :address, :phoneNumber, :role, :picture)
	RETURNING *
	`
	result := UserForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func UpdateUser(data UserForm) (UserForm, error) {
	sql := `UPDATE "users" SET
	"fullName"=COALESCE(NULLIF(:fullName, ''),"fullName"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"password"=COALESCE(NULLIF(:password, ''),"password"),
	"address"=COALESCE(NULLIF(:address, ''),"address"),
	"picture"=COALESCE(NULLIF(:picture, ''),"picture"),
	"phoneNumber"=COALESCE(NULLIF(:phoneNumber, ''),"phoneNumber"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := UserForm{}
	rows, err := db.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		rows.StructScan(&result)
	}

	return result, err
}

func DeleteUser(id int) (UserForm, error) {
	sql := `DELETE FROM "users" WHERE id = $1 RETURNING *`
	data := UserForm{}
	err := db.Get(&data, sql, id)
	return data, err
}

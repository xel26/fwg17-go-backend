package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = lib.DB

type User struct {
	Id          int          `db:"id" json:"id"`
	FullName    string       `db:"fullName" json:"fullName" form:"fullName"`
	Email       string       `db:"email" json:"email" form:"email"`
	Password    string       `db:"password" json:"password" form:"password"`
	Address     sql.NullString       `db:"address" json:"address" form:"address"`
	Picture     sql.NullString       `db:"picture" json:"picture"`
	PhoneNumber sql.NullString       `db:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        string       `db:"role" json:"role" form:"role"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `db:"updatedAt" json:"updatedAt"`
}

type Info struct{
	Data []User
	Count int
}

func FindAllUsers(searchKey string, sortBy string, order string, limit int, offset int) (Info, error) {
	sql := `
	SELECT * FROM "users" 
	WHERE "fullName" ILIKE $1
	ORDER BY $2 $3
	LIMIT $4 OFFSET $5
	`
	sqlCount := `
	SELECT COUNT(*) FROM "users"
	WHERE "fullName" ILIKE $1
	`

	result := Info{}
	data := []User{}
	err := db.Select(&data, sql, searchKey, sortBy, order, limit, offset)
	result.Data = data
	
	row := db.QueryRow(sqlCount, searchKey)
	err = row.Scan(&result.Count)

	return result, err
}


func FindOneUsers(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE id = $1`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}



// runtime error: invalid memory address or nil pointer dereference for rows.Next()
func CreateUser(data User) (User, error) {
	sql := `INSERT INTO "users" ("fullName", "email", "password", "address", "phoneNumber", "role") 
	VALUES
	(:fullName, :email, :password, :address, :phoneNumber, :role)
	RETURNING *
	`
	result := User{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next(){
		rows.StructScan(&result)
	}

	return result, err
}



// runtime error: invalid memory address or nil pointer dereference for rows.Next()
func UpdateUser(data User) (User, error) {
	sql := `UPDATE "users" SET
	fullName=:fullName,
	email=COALESCE(NULLIF(:email, ''),email),
	password=COALESCE(NULLIF(:password, ''),password),
	address=COALESCE(NULLIF(:address, ''),address),
	phoneNumber=COALESCE(NULLIF(:phoneNumber, ''),phoneNumber),
	role=COALESCE(NULLIF(:role, ''),role)
	WHERE id=:id
	RETURNING *
	`
	result := User{}
	rows, err := db.NamedQuery(sql, data)

	for rows.Next(){
		rows.StructScan(&result)
	}

	return result, err
}



func DeleteUser(id int) (User, error) {
	sql := `DELETE FROM "users" WHERE id = $1 RETURNING *`
	data := User{}
	err := db.Get(&data, sql, id)
	return data, err
}
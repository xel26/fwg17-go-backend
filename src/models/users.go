package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbU *sqlx.DB = lib.DB

type User struct {
	Id          int          `dbU:"id" json:"id"`
	FullName    string       `dbU:"fullName" json:"fullName" form:"fullName"`
	Email       string       `dbU:"email" json:"email" form:"email"`
	Password    string       `dbU:"password" json:"password" form:"password"`
	Address     sql.NullString       `dbU:"address" json:"address" form:"address"`
	Picture     sql.NullString       `dbU:"picture" json:"picture"`
	PhoneNumber sql.NullString       `dbU:"phoneNumber" json:"phoneNumber" form:"phoneNumber"`
	Role        string       `dbU:"role" json:"role" form:"role"`
	CreatedAt   time.Time    `dbU:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `dbU:"updatedAt" json:"updatedAt"`
}

type Info struct{
	Data []User
	Count int
}

func FindAllUsers(searchKey string, sortBy string, order string, limit int, offset int) (Info, error) {
	sql := `
	SELECT * FROM "users" 
	WHERE "fullName" ILIKE $1
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $2 OFFSET $3
	`
	sqlCount := `
	SELECT COUNT(*) FROM "users"
	WHERE "fullName" ILIKE $1
	`

	result := Info{}
	data := []User{}
	err := dbU.Select(&data, sql,"%"+searchKey+"%", limit, offset)
	result.Data = data
	
	row := dbU.QueryRow(sqlCount, "%"+searchKey+"%")
	err = row.Scan(&result.Count)

	return result, err
}


func FindOneUsers(id int) (User, error) {
	sql := `SELECT * FROM "users" WHERE id = $1`
	data := User{}
	err := dbU.Get(&data, sql, id)
	return data, err
}




func CreateUser(data User) (User, error) {
	sql := `INSERT INTO "users" ("fullName", "email", "password", "address", "phoneNumber", "role") 
	VALUES
	(:fullName, :email, :password, :address, :phoneNumber, :role)
	RETURNING *
	`
	result := User{}
	rows, err := dbU.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}
	
	for rows.Next(){
		rows.StructScan(&result)
	}
	
	return result, err
}




func UpdateUser(data User) (User, error) {
	sql := `UPDATE "users" SET
	"fullName"=COALESCE(NULLIF(:fullName, ''),"fullName"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"password"=COALESCE(NULLIF(:password, ''),"password"),
	"address"=COALESCE(NULLIF(:address, ''),"address"),
	"phoneNumber"=COALESCE(NULLIF(:phoneNumber, ''),"phoneNumber"),
	"role"=COALESCE(NULLIF(:role, ''),"role")
	WHERE id=:id
	RETURNING *
	`
	result := User{}
	rows, err := dbU.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next(){
		rows.StructScan(&result)
	}

	return result, err
}



func DeleteUser(id int) (User, error) {
	sql := `DELETE FROM "users" WHERE id = $1 RETURNING *`
	data := User{}
	err := dbU.Get(&data, sql, id)
	return data, err
}
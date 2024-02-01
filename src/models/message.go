package models

import (
	"coffe-shop-be-golang/src/lib"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var dbM *sqlx.DB = lib.DB

type Message struct {
	Id          int          `dbM:"id" json:"id"`
	RecipientId int          `dbM:"recipientId" json:"recipientId" form:"recipientId"`
	SenderId    int          `dbM:"senderId" json:"senderId" form:"senderId"`
	Text        string       `dbM:"text" json:"text" form:"text"`
	CreatedAt   time.Time    `dbM:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `dbM:"updatedAt" json:"updatedAt"`
}

type InfoM struct {
	Data  []Message
	Count int
}



func FindAllMessage(sortBy string, order string, limit int, offset int) (InfoM, error) {
	sql := `
	SELECT * FROM "message" 
	ORDER BY "`+sortBy+`" `+order+`
	LIMIT $1 OFFSET $2
	`
	sqlCount := `
	SELECT COUNT(*) FROM "message"`

	result := InfoM{}
	data := []Message{}
	err := dbM.Select(&data, sql, limit, offset)
	result.Data = data
	
	row := dbM.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}



func FindOneMessage(id int) (Message, error) {
	sql := `SELECT * FROM "message" WHERE id = $1`
	data := Message{}
	err := dbM.Get(&data, sql, id)
	return data, err
}



func CreateMessage(data Message) (Message, error) {
	sql := `INSERT INTO "message" ("recipientId", "senderId", "text") 
	VALUES
	(:recipientId, :senderId, :text)
	RETURNING *
	`
	result := Message{}
	rows, err := dbM.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}
	
	for rows.Next(){
		rows.StructScan(&result)
	}
	
	return result, err
}




func UpdateMessage(data Message) (Message, error) {
	sql := `UPDATE "message" SET
	"recipientId"=COALESCE(NULLIF(:recipientId, ''),"recipientId"),
	"senderId"=COALESCE(NULLIF(:senderId, ''),"senderId"),
	"text"=COALESCE(NULLIF(:text, ''),"text")
	WHERE id=:id
	RETURNING *
	`
	result := Message{}
	rows, err := dbM.NamedQuery(sql, data)
	if err != nil {
		return result, err
	}

	for rows.Next(){
		rows.StructScan(&result)
	}

	return result, err
}



func DeleteMessage(id int) (Message, error) {
	sql := `DELETE FROM "message" WHERE id = $1 RETURNING *`
	data := Message{}
	err := dbM.Get(&data, sql, id)
	return data, err
}
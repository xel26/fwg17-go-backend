package models

import (
	"database/sql"
	"time"
)

type Message struct {
	Id          int          `db:"id" json:"id"`
	RecipientId int          `db:"recipientId" json:"recipientId" form:"recipientId" binding:"required,numeric"`
	SenderId    int          `db:"senderId" json:"senderId" form:"senderId" binding:"required,numeric"`
	Text        string       `db:"text" json:"text" form:"text" binding:"required"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   sql.NullTime `db:"updatedAt" json:"updatedAt"`
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
	err := db.Select(&data, sql, limit, offset)
	result.Data = data
	
	row := db.QueryRow(sqlCount)
	err = row.Scan(&result.Count)

	return result, err
}



func FindOneMessage(id int) (Message, error) {
	sql := `SELECT * FROM "message" WHERE id = $1`
	data := Message{}
	err := db.Get(&data, sql, id)
	return data, err
}



func CreateMessage(data Message) (Message, error) {
	sql := `INSERT INTO "message" ("recipientId", "senderId", "text") 
	VALUES
	(:recipientId, :senderId, :text)
	RETURNING *
	`
	result := Message{}
	rows, err := db.NamedQuery(sql, data)
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
	"recipientId"=COALESCE(NULLIF(:recipientId, 0),"recipientId"),
	"senderId"=COALESCE(NULLIF(:senderId, 0),"senderId"),
	"text"=COALESCE(NULLIF(:text, ''),"text"),
	"updatedAt"=NOW()
	WHERE id=:id
	RETURNING *
	`
	result := Message{}
	rows, err := db.NamedQuery(sql, data)
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
	err := db.Get(&data, sql, id)
	return data, err
}
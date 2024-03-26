package models

import (
	"time"

	"github.com/LukaGiorgadze/gonull"
)

type intRandom struct {
	Id        int           `db:"id" json:"id"`
	IntRand       string        `db:"intRand" json:"intRand"`
	CreatedAt time.Time     `db:"createdAt" json:"createdAt"`
	UpdatedAt gonull.Nullable[time.Time] `db:"updatedAt" json:"updatedAt"`
}




func CreateIntRandom(intRand string) (intRandom, error) {
	sql := `INSERT INTO "intRandom" ("intRand") 
	VALUES ($1) 
	RETURNING *
	`
	result := intRandom{}
	row := db.QueryRow(sql, intRand)
	err := row.Scan(&result)

	return result, err
}


func FindOneByIntRandom(intRand string) (intRandom, error) {
	sql := `SELECT * FROM "intRandom" WHERE "intRand" = $1`
	data := intRandom{}
	err := db.Get(&data, sql, intRand)
	return data, err
}

func DeleteIntRandom(intRand string) (intRandom, error) {
	sql := `DELETE FROM "intRandom" WHERE "intRand" = $1 RETURNING *`
	data := intRandom{}
	err := db.Get(&data, sql, intRand)
	return data, err
}

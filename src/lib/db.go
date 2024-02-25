package lib

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func connectDB()*sqlx.DB{
	// db, err := sqlx.Connect("postgres", `user=postgres dbname=`+os.Getenv("DB_NAME")+` password=`+os.Getenv("DB_PASSWORD")+` port=`+os.Getenv("DB_PORT")+` host=`+os.Getenv("DB_HOST")+` sslmode=disable`) // error!
	db, err := sqlx.Connect("postgres", `user=postgres dbname=db_coffee_shop password=1 port=5433 host=143.110.156.215 sslmode=disable`)
	if err !=  nil {
		fmt.Println(err)
	}
	return db
}

var DB *sqlx.DB = connectDB()
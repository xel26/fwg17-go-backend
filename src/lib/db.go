package lib

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func connectDB()*sqlx.DB{
	// db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
    // os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT")))

	db, err := sqlx.Connect("postgres", `user=postgres.ircpdmthfidwfvchivrw password=5uJoOUiAbUl57U7X host=aws-0-ap-southeast-1.pooler.supabase.com port=5432 dbname=postgres`)

	// db, err := sqlx.Connect("postgres", `user=postgres password=1 host=143.110.156.215 port=5433 dbname=postgres`)


	if err !=  nil {
		fmt.Println(err)
	}
	return db
}

var DB *sqlx.DB = connectDB()
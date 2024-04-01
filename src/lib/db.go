package lib

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func connectDB()*sqlx.DB{
	// //dimatikan saat build image
	// err := godotenv.Load()
    // if err != nil {
    //     fmt.Println("Error loading .env file")
    // }
	
	dbConnect := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
    os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	db, err := sqlx.Connect("postgres", dbConnect)

	// db, err := sqlx.Connect("postgres", `user=postgres password=1 host=143.110.156.215 port=5433 dbname=postgres`)


	if err !=  nil {
		fmt.Println(err)
	}
	return db
}

var DB *sqlx.DB = connectDB()
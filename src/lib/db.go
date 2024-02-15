package lib

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)


func connectDB()*sqlx.DB{
	// db, err := sqlx.Connect("postgres", "user=postgres dbname=coffee_shop_implementation_basic password=1 sslmode=disable")
	db, err := sqlx.Connect("postgres", `user=postgres dbname=volume_data_coffee_shop password=`+os.Getenv("DB_PASSWORD")+` port=5444 host=host.docker.internal sslmode=disable`)
	if err !=  nil {
		fmt.Println(err)
	}
	return db
}

var DB *sqlx.DB = connectDB()
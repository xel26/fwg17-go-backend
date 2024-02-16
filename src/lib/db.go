package lib

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// POSTGRES_URL="postgresql://postgres:1@db:5432/volume_data_coffee_shop"

func connectDB()*sqlx.DB{
	// db, err := sqlx.Connect("postgres", "user=postgres dbname=coffee_shop_implementation_basic password=1 sslmode=disable")
	db, err := sqlx.Connect("postgres", `user=postgres dbname=volume_data_coffee_shop password=1 port=5444 host=host.docker.internal sslmode=disable`)
	// db, err := sqlx.Connect("postgres", `user=postgres dbname=volume_data_coffee_shop password=`+os.Getenv("DB_PASSWORD")+` port=5444 host=`+os.Getenv("DB_HOST")+` sslmode=disable`)
	// db, err := sqlx.Connect("postgres", `user=postgres dbname=volume_data_coffee_shop password=`+os.Getenv("DB_PASSWORD")+` port=5444 host=db sslmode=disable`)
	if err !=  nil {
		fmt.Println(err)
	}
	return db
}

var DB *sqlx.DB = connectDB()
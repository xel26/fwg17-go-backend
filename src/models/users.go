package models

import (
	"database/sql"
	"log"
	"sync"

	"github.com/xel26/fwg17-go-backend/src/lib"
)

func GetAllUsers() (*sql.Rows, error) {
	var wg sync.WaitGroup

	db, err := lib.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
	  panic(err)
	}

	wg.Add(1)

	var rows *sql.Rows
	var errorl error

	go func() {
		wg.Done()
		rows, errorl = db.Query("SELECT * FROM users")
		if errorl != nil {
			panic(errorl.Error())
		}
		defer rows.Close()
	}()

	wg.Wait()
	return rows, err
}

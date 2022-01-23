package database

import (
	_"github.com/lib/pq"

	 "database/sql"
	 _"fmt"
	 "log"
)
type Db struct {
	*sql.DB
}

func (db *Db)OpenDb(){
	connStr := "user=postgres dbname=postgres password=postgres sslmode=disable"
	data, err := sql.Open("postgres", connStr)
	db.DB = data
	if err != nil {
		log.Fatal(err)
	}
}

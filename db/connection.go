package db 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// MysqlConnection is a function that returns a pointer to a sql.DB object

func Connection() *sql.DB {

	dataSourceName := "root:naog7412@(localhost:3306)/act2?parseTime=true"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	return db
	
}


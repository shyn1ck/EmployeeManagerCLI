package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var dbConn *sql.DB

func ConnectToDB() error {
	connStr := "user=postgres password=2003 dbname=employee_manager_cli_db sslmode=disable"
	var err error
	dbConn, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = dbConn.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to database.")
	return nil
}

func GetDBConn() *sql.DB {
	return dbConn
}

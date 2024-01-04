// database/db.go
package database

import (
	"fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

const (
  host     = "159.89.34.98"
  port     = 5432
  user     = "postgres"
  password = "Sharingan@1"
  dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Initialize your database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
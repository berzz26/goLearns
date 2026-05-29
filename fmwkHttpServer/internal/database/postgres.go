package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"

	//open a connection with drive as postgres and passing the connstr
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	//test the connection by pinging the db
	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging the db: ", err)
	}
	log.Println("Connected to the db")

	return db
}

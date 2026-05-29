package repository

import (
	"database/sql"

	"log"
)


func GetUser(db *sql.DB) {
	rows, err := db.Query(`
		SELECT id, name, email
		FROM users
	`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var name string
		var email string

		err := rows.Scan(
			&id,
			&name,
			&email,
		)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(id, name, email)
	}
}

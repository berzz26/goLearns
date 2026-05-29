package repository

import (
	"database/sql"
	"fmwkHttpServer/internal/models"
	"log"
)

func GetUser(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query(`
		SELECT id, name, email
		FROM users
	`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}
	return users, nil
}

func GetOneUser(
	db *sql.DB,
	id int,
) (*model.User, error) {

	var user model.User

	err := db.QueryRow(`
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
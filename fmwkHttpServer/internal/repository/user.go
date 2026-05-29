package repository

import (
	"database/sql"
	"fmwkHttpServer/internal/models"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(
	db *sql.DB,
) *UserRepository {

	return &UserRepository{
		db: db,
	}
}

// this is a struct method..can be called by UserRepository.GetUsers()
func (r *UserRepository) GetUsers() ([]model.User, error) {
	// query() used when expecting multiple results
	rows, err := r.db.Query(`
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

func (r *UserRepository) GetOneUser(id int) (*model.User, error) {

	var user model.User

	// queryrow used to get 1 result
	err := r.db.QueryRow(`
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

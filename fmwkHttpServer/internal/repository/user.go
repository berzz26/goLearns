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
		SELECT id, username
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
			&user.Username,
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
		SELECT id, username
		FROM users
		WHERE id = $1
	`, id).Scan(
		&user.ID,
		&user.Username,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) AddUser(user model.User) (*model.User, error) {
	err := r.db.QueryRow(
		`INSERT INTO users(id, username)
         VALUES($1, $2)
         RETURNING id`,
		user.ID,
		user.Username,
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

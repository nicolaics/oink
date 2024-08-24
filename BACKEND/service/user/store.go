package user

import (
	"database/sql"
	"fmt"
	_ "log"

	"github.com/nicolaics/oink/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		return nil, err
	}

	user := new(types.User)

	for rows.Next() {
		user, err = scanRowIntoUser(rows)

		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
		
	if err != nil {
		return nil, err
	}

	user := new(types.User)

	for rows.Next() {
		user, err = scanRowIntoUser(rows)

		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	
	return user, nil
}

func (s *Store) CreateUser(user types.User) (int, error) {
	_, err := s.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
						user.Name, user.Email, user.Password)
	if err != nil {
		return -1, err
	}

	u, err := s.GetUserByEmail(user.Email)
	if err != nil {
		return -1, err
	}

	return u.ID, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

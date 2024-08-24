package account

import (
	"database/sql"
	"fmt"

	"github.com/nicolaics/oink/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAccountByID(id int) (*types.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	account := new(types.Account)

	for rows.Next() {
		account, err = scanRowIntoAccount(rows)

		if err != nil {
			return nil, err
		}
	}

	if account.ID == 0 {
		return nil, fmt.Errorf("account not found")
	}

	return account, nil
}

func (s *Store) UpdateBalanceAmount(userId int, newBalance float64) error {
	_, err := s.db.Exec("UPDATE account JOIN users ON account.user_id = users.id SET balance = ? WHERE users.id = ? ",
							newBalance, userId)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *Store) CreateAccount(id int) error {
	_, err := s.db.Exec("INSERT INTO account (user_id) VALUES (?)", id)
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoAccount(rows *sql.Rows) (*types.Account, error) {
	account := new(types.Account)

	err := rows.Scan(
		&account.ID,
		&account.UserID,
		&account.Balance,
	)

	if err != nil {
		return nil, err
	}

	return account, nil
}

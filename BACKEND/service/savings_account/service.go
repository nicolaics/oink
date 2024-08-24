package savingsaccount

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

func (s *Store) GetSavingsAccountByID(id int) (*types.SavingsAccount, error) {
	rows, err := s.db.Query("SELECT * FROM savings_account WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	savingsAcc := new(types.SavingsAccount)

	for rows.Next() {
		savingsAcc, err = scanRowIntoSavingsAccount(rows)

		if err != nil {
			return nil, err
		}
	}

	if savingsAcc.ID == 0 {
		return nil, fmt.Errorf("savings account not found")
	}

	return savingsAcc, nil
}

func scanRowIntoSavingsAccount(rows *sql.Rows) (*types.SavingsAccount, error) {
	savingsAcc := new(types.SavingsAccount)

	err := rows.Scan(
		&savingsAcc.ID,
		&savingsAcc.UserID,
		&savingsAcc.Amount,
	)

	if err != nil {
		return nil, err
	}

	return savingsAcc, nil
}

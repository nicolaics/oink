package account

import (
	"database/sql"
	"fmt"

	"github.com/nicolaics/oink/types"
	"github.com/nicolaics/oink/utils"
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
		account, err = utils.ScanRowIntoAccount(rows)

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

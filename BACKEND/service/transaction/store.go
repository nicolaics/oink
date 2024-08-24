package transaction

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

func (s *Store) GetAccountByID(userId int) (*types.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE user_id = ? ", userId)

	if err != nil {
		return nil, err
	}

	acc := new(types.Account)

	for rows.Next() {
		acc, err = utils.ScanRowIntoAccount(rows)

		if err != nil {
			return nil, err
		}
	}

	if acc.ID == 0 {
		return nil, fmt.Errorf("account not found")
	}

	return acc, nil
}

func (s *Store) UpdateBalanceAmount(userId int, newBalance float64) error {
	_, err := s.db.Exec("UPDATE account JOIN users ON account.user_id = users.id SET balance = ? WHERE users.id = ? ",
							newBalance, userId)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *Store) GetTransactionsByID(userId int, reqType string) ([]types.Transaction, error) {
	rows := new(sql.Rows)
	var err error

	if reqType == "sender" {
		rows, err = s.db.Query("SELECT * FROM transaction WHERE sender_id = ? ", userId)
	} else {
		rows, err = s.db.Query("SELECT * FROM transaction WHERE receiver_id = ? ", userId)
	}
	if err != nil {
		return nil, err
	}

	transactions := make([]types.Transaction, 0)

	for rows.Next() {
		tx, err := scanRowIntoTransactions(rows)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, *tx)
	}

	return transactions, nil
}

func (s *Store) CreateTransaction(tx types.Transaction) error {
	_, err := s.db.Exec("INSERT INTO transaction (receiver_id, sender_id, amount) VALUES (?, ?, ?)",
						tx.ReceiverID, tx.SenderID, tx.Amount)
		
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoTransactions(rows *sql.Rows) (*types.Transaction, error) {
	tx := new(types.Transaction)

	err := rows.Scan(
		&tx.ID,
		&tx.ReceiverID,
		&tx.SenderID,
		&tx.Amount,
		&tx.TransactionTime,
	)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

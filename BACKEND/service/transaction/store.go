package transaction

import (
	"database/sql"

	"github.com/nicolaics/oink/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) UpdateBalanceAmount(userId int, newBalance float64) error {
	_, err := s.db.Exec("UPDATE account JOIN users ON account.user_id = users.id SET balance = ? WHERE users.id = ? ",
							newBalance, userId)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *Store) GetTransactionsByID(userId int) ([]types.Transaction, error) {
	rows, err := s.db.Query("SELECT * FROM transaction WHERE user_id = ? ", userId)
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
	_, err := s.db.Exec("INSERT INTO transaction (user_id, amount) VALUES (?, ?)",
						tx.UserID, tx.Amount)
		
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoTransactions(rows *sql.Rows) (*types.Transaction, error) {
	tx := new(types.Transaction)

	err := rows.Scan(
		&tx.ID,
		&tx.UserID,
		&tx.Amount,
		&tx.TransactionTime,
	)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

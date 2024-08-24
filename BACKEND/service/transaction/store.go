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

func (s *Store) GetTransactionsByID(userId int) ([]types.Transaction, error) {
	rows, err := s.db.Query("SELECT * FROM transaction WHERE user_id = ? AND visibile = ?", userId, true)
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
	_, err := s.db.Exec("INSERT INTO transaction (user_id, amount, src_acc, dest_acc, visible) VALUES (?, ?, ?, ?, ?)",
						tx.UserID, tx.Amount, tx.SrcAccount, tx.DestAccount, tx.Visible)
		
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateTransactionsVisibility(userId int) error {
	_, err := s.db.Exec("UPDATE transaction JOIN users ON transaction.user_id = users.id SET visible = ? WHERE visible = ? AND users.id = ?",
							false, true, userId)
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

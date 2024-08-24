package loan

import (
	"database/sql"
	// "fmt"

	"github.com/nicolaics/oink/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetLoansDataByDebtorID(userId int) ([]types.Loan, error) {
	rows, err := s.db.Query("SELECT * FROM transaction WHERE sender_id = ? ", userId)
	if err != nil {
		return nil, err
	}

	loans := make([]types.Loan, 0)

	for rows.Next() {
		loan, err := scanRowIntoLoan(rows)

		if err != nil {
			return nil, err
		}

		loans = append(loans, *loan)
	}

	return loans, nil
}

func (s *Store) CreateLoan(loan types.Loan) error {
	_, err := s.db.Exec("INSERT INTO loan (debtor_id, amount, start_date, end_date, duration) VALUES (?, ?, ?, ?, ?)",
						loan.DebtorID, loan.Amount, loan.StartDate, loan.EndDate, loan.Duration)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateLoanPayment(userId int, amountPaid float64) error {
	_, err := s.db.Exec("UPDATE loan JOIN users ON loan.debtor_id = users.id SET amount_paid = ? WHERE users.id = ? ",
							amountPaid, userId)
	if err != nil {
		return err
	}
	
	return nil
}

func scanRowIntoLoan(rows *sql.Rows) (*types.Loan, error) {
	loan := new(types.Loan)

	err := rows.Scan(
		&loan.ID,
		&loan.DebtorID,
		&loan.Amount,
		&loan.StartDate,
		&loan.EndDate,
		&loan.Duration,
	)

	if err != nil {
		return nil, err
	}

	return loan, nil
}

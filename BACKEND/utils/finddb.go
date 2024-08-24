package utils

import (
	"database/sql"

	"github.com/nicolaics/oink/types"
)

func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func ScanRowIntoAccount(rows *sql.Rows) (*types.Account, error) {
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
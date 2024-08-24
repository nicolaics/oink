package types

import "time"

type RegisterUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}
type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(string) (*User, error)
	GetUserByID(int) (*User, error)
	CreateUser(User) (int, error)
	CreateAccount(int) error
	CreateSavingsAccount(int) error
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"string"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewTransactionPayload struct {
	ReceiverID      int       `json:"receiverId" validate:"required"`
	SenderID        int       `json:"senderId" validate:"required"`
	Amount          float64   `json:"amount" validate:"required"`
	TransactionTime time.Time `json:"txTime" validate:"required"`
}

type TransactionPayload struct {
	UserID      int       `json:"userId" validate:"required"`
}


type TransactionStore interface {
	GetTransactionsByID(int, string) ([]Transaction, error)
	CreateTransaction(Transaction) error
	UpdateBalanceAmount(userId int, newBalance float64) error
	GetAccountByID(int) (*Account, error)
}

type Transaction struct {
	ID              int       `json:"id"`
	ReceiverID      int       `json:"receiverId"`
	SenderID        int       `json:"senderId"`
	Amount          float64   `json:"amount"`
	TransactionTime time.Time `json:"txTime"`
}

type LoanStore interface {
	GetLoanDataByDebtorID(int) (*Loan, error)
	CreateLoan(Loan) error
}

type Loan struct {
	ID        int       `json:"id"`
	DebtorID  int       `json:"debtorId"`
	Amount    float64   `json:"amount"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Duration  string    `json:"duration"`
}

type AccountStore interface {
	GetAccountByID(int) (*Account, error)
	UpdateBalanceAmount(userId int, newBalance float64) error
}

type AccountPayload struct {
	UserID  int     `json:"userId" validate:"required"`
	Balance float64 `json:"balance"`
}

type Account struct {
	ID      int     `json:"id"`
	UserID  int     `json:"userId"`
	Balance float64 `json:"balance"`
}

type SavingsAccountStore interface {
	GetSavingsAccountByID(int) (*SavingsAccount, error)
	UpdateSavingsAmount(userId int, amount float64) error
}

type SavingsAmountPayload struct {
	UserID    int     `json:"userId" validate:"required"`
	NewAmount float64 `json:"newAmount"`
}

type SavingsAccount struct {
	ID     int     `json:"id"`
	UserID int     `json:"userId"`
	Amount float64 `json:"amount"`
}

type PigRaceStore interface {
	GetPigStaminaByID(int) (*PigRace, error)
	UpdateFinalDistance(userId int, distance float64) error
	UpdatePigStamina(userId int, stamina float64) error
}

type PigRace struct {
	ID                  int     `json:"id"`
	UserID              int     `json:"userId"`
	PigStamina          float64 `json:"pigStamina"`
	FinalDistanceToGoal float64 `json:"finalDistanceToGoal"`
}

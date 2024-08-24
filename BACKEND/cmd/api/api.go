package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nicolaics/oink/service/account"
	"github.com/nicolaics/oink/service/loan"
	savingsaccount "github.com/nicolaics/oink/service/savings_account"
	"github.com/nicolaics/oink/service/transaction"
	"github.com/nicolaics/oink/service/user"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	accountStore := account.NewStore(s.db)
	accountHandler := account.NewHandler(accountStore)
	accountHandler.RegisterRoutes(subrouter)
	
	savingsAccountStore := savingsaccount.NewStore(s.db)
	savingsAccountHandler := savingsaccount.NewHandler(savingsAccountStore)
	savingsAccountHandler.RegisterRoutes(subrouter)

	transactionStore := transaction.NewStore(s.db)
	transactionHandler := transaction.NewHandler(transactionStore)
	transactionHandler.RegisterRoutes(subrouter)

	loanStore := loan.NewStore(s.db)
	loanHandler := loan.NewHandler(loanStore, userStore)
	loanHandler.RegisterRoutes(subrouter)

	log.Println("Listening on: ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
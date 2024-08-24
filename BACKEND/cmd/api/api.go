package api

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/nicolaics/oink/service/account"
	"github.com/nicolaics/oink/service/loan"
	"github.com/nicolaics/oink/service/pigrace"
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

type LogResponseWriter struct {
    http.ResponseWriter
    statusCode int
    buf        bytes.Buffer
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
    return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
    w.statusCode = code
    w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
    w.buf.Write(body)
    return w.ResponseWriter.Write(body)
}

type LogMiddleware struct {
    logger *log.Logger
}

func NewLogMiddleware(logger *log.Logger) *LogMiddleware {
    return &LogMiddleware{logger: logger}
}

func (m *LogMiddleware) Func() mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            startTime := time.Now()

            logRespWriter := NewLogResponseWriter(w)
            next.ServeHTTP(logRespWriter, r)

            m.logger.Printf(
                "duration=%s status=%d body=%s",
                time.Since(startTime).String(),
                logRespWriter.statusCode,
                logRespWriter.buf.String())
        })
    }
}

func (s *APIServer) Run() error {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	accountStore := account.NewStore(s.db)
	accountHandler := account.NewHandler(accountStore)
	accountHandler.RegisterRoutes(subrouter)
	
	savingsAccountStore := savingsaccount.NewStore(s.db)
	savingsAccountHandler := savingsaccount.NewHandler(savingsAccountStore, accountStore)
	savingsAccountHandler.RegisterRoutes(subrouter)

	transactionStore := transaction.NewStore(s.db)
	transactionHandler := transaction.NewHandler(transactionStore, accountStore, savingsAccountStore)
	transactionHandler.RegisterRoutes(subrouter)

	pigRaceStore := pigrace.NewStore(s.db)
	pigRaceHandler := pigrace.NewHandler(pigRaceStore)
	pigRaceHandler.RegisterRoutes(subrouter)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore, accountStore, savingsAccountStore, pigRaceStore)
	userHandler.RegisterRoutes(subrouter)

	loanStore := loan.NewStore(s.db)
	loanHandler := loan.NewHandler(loanStore, userStore)
	loanHandler.RegisterRoutes(subrouter)
	
	log.Println("Listening on: ", s.addr)

	logMiddleware := NewLogMiddleware(logger)
    router.Use(logMiddleware.Func())

	return http.ListenAndServe(s.addr, router)
}
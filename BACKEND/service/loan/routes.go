package loan

import (
	"fmt"
	"go/types"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nicolaics/oink/types"
	"github.com/nicolaics/oink/utils"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type Handler struct {
	loanStore types.LoanStore
	userStore types.UserStore
}

func NewHandler(loanStore types.LoanStore, userStore types.UserStore) *Handler {
	return &Handler{loanStore: loanStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/loan", h.handleGetLoanData).Methods(http.MethodGet)
	router.HandleFunc("/loan", h.handleNewLoan).Methods(http.MethodPost)
	router.HandleFunc("/loan", h.handleLoanPayment).Methods(http.MethodPatch)
}

func (h *Handler) handleGetLoanData(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.LoanPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	loans, err := h.loanStore.GetLoansDataByDebtorID(payload.DebtorID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, loans)
}


func (h *Handler) handleNewLoan(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.LoanPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.userStore.GetUserByID(payload.DebtorID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	err = h.loanStore.CreateLoan(types.Loan{
		DebtorID: payload.DebtorID,
		Amount: payload.Amount,
		StartDate: payload.StartDate,
		EndDate: payload.EndDate,
		Duration: payload.Duration,
		Active: true,

	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h * Handler) handleLoanPayment(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.LoanPaymentPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	loans, err := h.userStore.GetLoansDataByDebtorID(payload.DebtorID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	var loan types.Loan

	for _, l := range(loans) {
		if l.active {
			loan = l
			break
		}
	}

	err = h.loanStore.UpdateLoanPayment(payload.DebtorID, (payload.PaymentAmount + loan.AmountPaid))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}


package transaction

import (
	"fmt"
	"math"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nicolaics/oink/types"
	"github.com/nicolaics/oink/utils"
)

type Handler struct {
	transactionStore   types.TransactionStore
	accountStore       types.AccountStore
	savingAccountStore types.SavingsAccountStore
}

func NewHandler(transactionStore types.TransactionStore, accountStore types.AccountStore,
				savingAccountStore types.SavingsAccountStore) *Handler {
	return &Handler{
		transactionStore: transactionStore,
		accountStore: accountStore,
		savingAccountStore: savingAccountStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/transaction/create", h.handleCreateTransaction).Methods(http.MethodPost)
	router.HandleFunc("/transaction", h.handleGetTransactionByID).Methods(http.MethodPost)
	router.HandleFunc("/transaction/create", func(w http.ResponseWriter, r *http.Request) { utils.WriteJSONForOptions(w, http.StatusOK, nil) }).Methods(http.MethodOptions)
	router.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) { utils.WriteJSONForOptions(w, http.StatusOK, nil) }).Methods(http.MethodOptions)
}

func (h *Handler) handleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.NewTransactionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	account, err := h.accountStore.GetAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	factor := math.Pow(10, float64(3)-1)
	roundedNumber := math.Ceil(payload.Amount/factor) * factor

	savingAmount := math.Abs(roundedNumber - payload.Amount)

	err = h.accountStore.UpdateBalanceAmount(payload.UserID, (account.Balance + roundedNumber))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	savingAcc, err := h.savingAccountStore.GetSavingsAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	err = h.savingAccountStore.UpdateSavingsAmount(payload.UserID, (savingAcc.Amount + savingAmount))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	err = h.transactionStore.CreateTransaction(types.Transaction{
		UserID: payload.UserID,
		Amount: roundedNumber,
		SrcAccount: "",
		DestAccount: types.ACCOUNT,
		Visible: true,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.transactionStore.CreateTransaction(types.Transaction{
		UserID: payload.UserID,
		Amount: savingAmount,
		SrcAccount: types.ACCOUNT,
		DestAccount: types.SAVINGS,
		Visible: false,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleGetTransactionByID(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.TransactionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	transactions, err := h.transactionStore.GetTransactionsByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, transactions)
}

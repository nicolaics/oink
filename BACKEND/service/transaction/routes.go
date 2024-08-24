package transaction

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nicolaics/oink/types"
	"github.com/nicolaics/oink/utils"
)

type Handler struct {
	store types.TransactionStore
}

func NewHandler(store types.TransactionStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/transaction", h.handleCreateTransaction).Methods(http.MethodPost)
	router.HandleFunc("/transaction/{reqType}", h.handleGetTransactionByID).Methods(http.MethodGet)

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

	senderAcc, err := h.store.GetAccountByID(payload.SenderID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid sender account ID"))
		return
	}

	receiverAcc, err := h.store.GetAccountByID(payload.ReceiverID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid receiver account ID"))
		return
	}

	err = h.store.UpdateBalanceAmount(payload.SenderID, (payload.Amount + senderAcc.Balance))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.UpdateBalanceAmount(payload.SenderID, (receiverAcc.Balance - payload.Amount))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.CreateTransaction(types.Transaction{
		ReceiverID: payload.ReceiverID,
		SenderID: payload.SenderID,
		Amount: payload.Amount,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, []types.Account{*senderAcc, *receiverAcc})
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

	vars := mux.Vars(r)
	reqType := vars["reqType"]

	transactions, err := h.store.GetTransactionsByID(payload.UserID, reqType)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, transactions)
}

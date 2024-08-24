package account

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nicolaics/oink/types"
	"github.com/nicolaics/oink/utils"
)

type Handler struct {
	accStore types.AccountStore
	txStore types.TransactionStore
}

func NewHandler(accStore types.AccountStore, txStore types.TransactionStore) *Handler {
	return &Handler{accStore: accStore, txStore: txStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/account/balance", h.handleUpdateBalance).Methods(http.MethodPatch)
	router.HandleFunc("/account/balance", h.handleGetBalanceAmount).Methods(http.MethodPost)
	router.HandleFunc("/account/balance", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSONForOptions(w, http.StatusOK, nil)}).Methods(http.MethodOptions)
}

// add cash from outside to this account only
func (h *Handler) handleUpdateBalance(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.AccountPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	acc, err := h.accStore.GetAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid account ID"))
		return
	}

	err = h.accStore.UpdateBalanceAmount(payload.UserID, (acc.Balance + payload.Balance))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.txStore.CreateTransaction(types.Transaction{
		UserID: payload.UserID,
		Amount: payload.Balance,
		SrcAccount: "",
		DestAccount: types.ACCOUNT,
		Visible: true,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleGetBalanceAmount(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.AccountPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	account, err := h.accStore.GetAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, account)
}

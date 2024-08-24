package savingsaccount

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nicolaics/oink/types"
	"github.com/nicolaics/oink/utils"
)

type Handler struct {
	savingAccStore types.SavingsAccountStore
	accStore types.AccountStore
	txStore types.TransactionStore
}

func NewHandler(savingAccStore types.SavingsAccountStore, accStore types.AccountStore, txStore types.TransactionStore) *Handler {
	return &Handler{savingAccStore: savingAccStore, accStore: accStore, txStore: txStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/savings-account", h.handleEmptySavingsAmount).Methods(http.MethodPatch)
	router.HandleFunc("/savings-account", h.handleGetSavingsAmount).Methods(http.MethodPost)
	router.HandleFunc("/savings-account", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSONForOptions(w, http.StatusOK, nil)}).Methods(http.MethodOptions)
}

func (h *Handler) handleEmptySavingsAmount(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.SavingsAmountPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	savingAcc, err := h.savingAccStore.GetSavingsAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid account ID"))
		return
	}

	acc, err := h.accStore.GetAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid account ID"))
		return
	}

	err = h.savingAccStore.UpdateSavingsAmount(payload.UserID, 0.0)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.accStore.UpdateBalanceAmount(payload.UserID, (acc.Balance + savingAcc.Amount))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.txStore.CreateTransaction(types.Transaction{
		UserID: payload.UserID,
		Amount: savingAcc.Amount,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.txStore.CreateTransaction(types.Transaction{
		UserID: payload.UserID,
		Amount: -(savingAcc.Amount),
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}



func (h *Handler) handleGetSavingsAmount(w http.ResponseWriter, r *http.Request) {
	log.Println(r)

	// get JSON Payload
	var payload types.SavingsAmountPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	log.Println(payload)
	
	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	savingsAcc, err := h.savingAccStore.GetSavingsAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, savingsAcc)
}

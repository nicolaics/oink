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
	store types.AccountStore
}

func NewHandler(store types.AccountStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/account/balance", h.handleUpdateBalance).Methods(http.MethodPatch)
	router.HandleFunc("/account/balance", h.handleGetBalanceAmount).Methods(http.MethodPost)
	router.HandleFunc("/account/balance", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSONForOptions(w, http.StatusOK, "options")}).Methods(http.MethodOptions)
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSON(w, http.StatusOK, "random")}).Methods(http.MethodGet)
	router.HandleFunc("/test_", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSON(w, http.StatusOK, "random_2")}).Methods(http.MethodPost)
}

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

	err := h.store.UpdateBalanceAmount(payload.UserID, payload.Balance)
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

	account, err := h.store.GetAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, account)
}

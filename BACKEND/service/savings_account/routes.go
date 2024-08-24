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
	store types.SavingsAccountStore
}

func NewHandler(store types.SavingsAccountStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/savings-account", h.handleUpdateSavingsAmount).Methods(http.MethodPatch)
	router.HandleFunc("/savings-account", h.handleGetSavingsAmount).Methods(http.MethodPost)
	router.HandleFunc("/savings-account", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSONForOptions(w, http.StatusOK, nil)}).Methods(http.MethodOptions)
}

func (h *Handler) handleUpdateSavingsAmount(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	// get JSON Payload
	var payload types.SavingsAmountPayload
	
	log.Println(payload)

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.UpdateSavingsAmount(payload.UserID, payload.NewAmount)
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

	savingsAcc, err := h.store.GetSavingsAccountByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, savingsAcc)
}

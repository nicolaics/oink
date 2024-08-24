package savingsaccount

import (
	"fmt"
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
	router.HandleFunc("/savings-account/savings", h.handleUpdateSavingsAmount).Methods(http.MethodPatch)
	// router.HandleFunc("/user/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleUpdateSavingsAmount(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.UpdateSavingsAmountPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

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

	err = h.store.UpdateSavingsAmount(savingsAcc, payload.NewAmount)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

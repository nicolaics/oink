package pigrace

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nicolaics/oink/types"
	"github.com/nicolaics/oink/utils"
)

type Handler struct {
	pigRaceStore types.PigRaceStore
}

func NewHandler(pigRaceStore types.PigRaceStore) *Handler {
	return &Handler{pigRaceStore: pigRaceStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/pigrace", h.handleGetPigRaceData).Methods(http.MethodPost)
	router.HandleFunc("/pigrace/stamina", h.handleUpdatePigStamina).Methods(http.MethodPatch)
	router.HandleFunc("/pigrace/distance", h.handleUpdateFinalDistance).Methods(http.MethodPatch)
}

func (h *Handler) handleGetPigRaceData(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.PigRacePayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	pigRace, err := h.pigRaceStore.GetPigRaceDataByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, pigRace)
}


func (h *Handler) handleUpdatePigStamina(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.PigRacePayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	pigRaceData, err := h.pigRaceStore.GetPigRaceDataByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	err = h.pigRaceStore.UpdatePigStamina(pigRaceData, payload.AdditionalStamina)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateFinalDistance(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.PigRacePayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	pigRaceData, err := h.pigRaceStore.GetPigRaceDataByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid account ID"))
		return
	}

	err = h.pigRaceStore.UpdateFinalDistance(pigRaceData, payload.AdditionalStamina)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

package user

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nicolaics/oink/config"
	"github.com/nicolaics/oink/service/auth"
	"github.com/nicolaics/oink/types"
	"github.com/nicolaics/oink/utils"
)

type Handler struct {
	userStore types.UserStore
	accountStore types.AccountStore
	savingsAccountStore types.SavingsAccountStore
}

func NewHandler(userStore types.UserStore, accountStore types.AccountStore,
				savingsAccountStore types.SavingsAccountStore) *Handler {
	return &Handler{
		userStore: userStore,
		accountStore: accountStore,
		savingsAccountStore: savingsAccountStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user/login", h.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSONForOptions(w, http.StatusOK, nil)}).Methods(http.MethodOptions)
	router.HandleFunc("/user/register", h.handleRegister).Methods(http.MethodPost)
	router.HandleFunc("/user/register", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSONForOptions(w, http.StatusOK, nil)}).Methods(http.MethodOptions)
	router.HandleFunc("/user/leaderboard", h.handleLeaderboard).Methods(http.MethodGet)
	router.HandleFunc("/user/leaderboard", func(w http.ResponseWriter, r *http.Request) {utils.WriteJSONForOptions(w, http.StatusOK, nil)}).Methods(http.MethodOptions)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.LoginUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	user, err := h.userStore.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no email found"))
		return
	}

	log.Printf("user Password: %s", user.Password)
	log.Printf("payload pw : %s", payload.Password)

	// check password match
	if !(auth.ComparePassword(user.Password, []byte(payload.Password))) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}


	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token, "userId": fmt.Sprintf("%d", user.ID), "name": user.Name})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON Payload
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// check if the user exists
	_, err := h.userStore.GetUserByEmail(payload.Email)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, 
						fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// if it doesn't, we create new user
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	userID, err := h.userStore.CreateUser(types.User{
		Name: payload.Name,
		Email: payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	err = h.accountStore.CreateAccount(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	
	err = h.savingsAccountStore.CreateSavingsAccount(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleLeaderboard(w http.ResponseWriter, r *http.Request) {
	// check if the user exists
	users, err := h.userStore.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	ranking := make([]float64, 0)

	for _, user := range(users) {
		saving, err := h.savingsAccountStore.GetSavingsAccountByID(user.ID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}

		ranking = append(ranking, saving.Amount)
	}

	slices.Sort(ranking)

	utils.WriteJSON(w, http.StatusCreated, ranking)
}

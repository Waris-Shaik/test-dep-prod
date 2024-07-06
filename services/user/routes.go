package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test-dep-prod/types"
	"test-dep-prod/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.handleRoot).Methods(http.MethodGet)
	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
	router.HandleFunc("/users", h.handleGetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", h.handleGetUser).Methods(http.MethodGet)

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start the route")
	var payload types.RegisterUserPayload
	// get the JSON payload from request body and parse it
	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Println("Error in parsing JSON:", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the JSON payload
	if err := utils.ValidateRegisterUserPayload(&payload); err != nil {
		log.Println("Error occured in validated payload:", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check the user already exists in db
	_, err := h.store.GetUserByEmail(payload.Email)

	if err == nil {
		log.Println("user with the credentials already exists")
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user already exists, please login"))
		return
	}

	// create the user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		UserName:  payload.UserName,
		Email:     payload.Email,
		Password:  payload.Password,
	})
	if err != nil {
		log.Println("Error creating user", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// return the success-response
	successResponse := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "user created successfully",
	}
	utils.WriteJSON(w, http.StatusCreated, successResponse)

}

func (h *Handler) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.store.GetAllUsers()
	if err != nil {
		log.Println("Error while retreiving all users", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return the success-response
	successResponse := struct {
		Success bool          `json:"success"`
		Users   []*types.User `json:"users"`
	}{
		Success: true,
		Users:   users,
	}
	utils.WriteJSON(w, http.StatusCreated, successResponse)

}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from request parameters
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		log.Println("UserID is required")
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("userID is required"))
		return
	}

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid user ID:", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user id"))
		return
	}

	// Retrieve user from the store
	user, err := h.store.GetUserByID(userID)
	if err != nil {
		log.Println("Error retrieving user:", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	// return the success response
	response := struct {
		Success bool       `json:"success"`
		User    types.User `json:"user"`
	}{
		Success: true,
		User:    *user,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleRoot(w http.ResponseWriter, r *http.Request) {
	// return the success-response
	successResponse := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "welcom to go-lang-server",
	}
	utils.WriteJSON(w, http.StatusCreated, successResponse)
}

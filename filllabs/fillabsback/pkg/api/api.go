package api

import (
	"encoding/json"
	"fillabs_intern_project/pkg/model"
	"fillabs_intern_project/pkg/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Api struct represents the API and holds a reference to the service.
type Api struct {
	service *service.Service
}

// NewApi creates a new instance of Api with the given service.
func NewApi(service *service.Service) *Api {
	return &Api{
		service: service,
	}
}

// AddUser handles the HTTP request to add a new user.
func (api Api) AddUser(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON request into a User model
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ReturnError(w, "error occurred when decoding body "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Add the user using the service
	addUser, err := api.service.AddUser(user)
	if err != nil {
		ReturnError(w, "error occurred when adding user "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response as JSON and send it back
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&addUser)
	if err != nil {
		ReturnError(w, "error occurred when encoding user "+err.Error(), http.StatusInternalServerError)
	}
}

// UpdateUser handles the HTTP request to update an existing user.
func (api Api) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, exists := params["user_id"]
	if !exists {
		ReturnError(w, "User ID not found ", http.StatusNotFound)
		return
	}

	// Convert the user ID from string to integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		ReturnError(w, "User ID can't be converted to integer "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user ID is a positive integer
	if userIDInt <= 0 {
		ReturnError(w, "User ID can't be equal to or less than zero "+err.Error(), http.StatusBadRequest)
		return
	}

	// Decode the incoming JSON request into a UserUpdateRequest model
	var userUpdateRequest model.UserUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&userUpdateRequest)
	if err != nil {
		ReturnError(w, "Update Request can't be decoded "+err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user using the service
	updatedUser, err := api.service.UpdateUser(uint(userIDInt), &userUpdateRequest)
	if err != nil {
		ReturnError(w, "User can't be updated "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response as JSON and send it back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser handles the HTTP request to delete an existing user.
func (api Api) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, exists := params["user_id"]
	if !exists {
		ReturnError(w, "User ID not found ", http.StatusNotFound)
		return
	}

	// Convert the user ID from string to integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		ReturnError(w, "User ID can't be converted to integer "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user ID is a positive integer
	if userIDInt <= 0 {
		ReturnError(w, "User ID can't be equal to or less than zero "+err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the user using the service
	err = api.service.DeleteUser(uint(userIDInt))
	if err != nil {
		ReturnError(w, "User can't be deleted "+err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare a success message and send it as a JSON response
	successfullyDeleted := map[string]string{
		"message": "User successfully deleted",
	}
	byteMessage, err := json.Marshal(successfullyDeleted)
	if err != nil {
		ReturnError(w, "User can't be deleted "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteMessage))
}

// GetUsers handles the HTTP request to retrieve a list of all users.
func (api Api) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Retrieve the list of users using the service
	users, err := api.service.GetUsers()
	if err != nil {
		ReturnError(w, "Users can't be listed "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response as JSON and send it back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

// GetUserByID handles the HTTP request to retrieve a user by its ID.
func (api Api) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, exists := params["user_id"]
	if !exists {
		ReturnError(w, "User ID not found ", http.StatusNotFound)
		return
	}

	// Convert the user ID from string to integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		ReturnError(w, "User ID can't be converted to integer "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user ID is a positive integer
	if userIDInt <= 0 {
		ReturnError(w, "User ID can't be equal to or less than zero "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the user by ID using the service
	user, err := api.service.GetUserByID(uint(userIDInt))
	if err != nil {
		ReturnError(w, "User can't be retrieved by ID "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response as JSON and send it back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

// ReturnError is a helper function to return an HTTP error with a specific message and status code.
func ReturnError(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}

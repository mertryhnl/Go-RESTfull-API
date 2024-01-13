package handler

import (
	"fillabs_intern_project/internal/storage"
	"fillabs_intern_project/pkg/api"
	"fillabs_intern_project/pkg/repository"
	"fillabs_intern_project/pkg/service"

	"github.com/gorilla/mux"
)

// InitHandler initializes the handler by setting up the necessary components,
// such as the database connection, repository, service, and API routes.
// It returns a Gorilla Mux router to be used in the application.
func InitHandler() (*mux.Router, error) {
	// Open a database connection
	db, err := storage.OpenDB()
	if err != nil {
		return nil, err
	}

	// Create a new repository with the database connection
	repo := repository.NewRepository(db)

	// Create a new service with the repository
	service := service.NewService(repo)

	// Create a new API instance with the service
	api := api.NewApi(service)

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define routes for various user-related operations
	router.HandleFunc("/users", api.GetUsers).Methods("GET")
	router.HandleFunc("/user/{user_id:[0-9]+}", api.GetUserByID).Methods("GET")
	router.HandleFunc("/users", api.AddUser).Methods("POST")
	router.HandleFunc("/users/{user_id:[0-9]+}", api.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{user_id:[0-9]+}", api.DeleteUser).Methods("DELETE")

	// Return the initialized router
	return router, nil
}

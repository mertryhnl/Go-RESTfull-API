package service

import (
	"errors"
	"fillabs_intern_project/pkg/model"
	"fillabs_intern_project/pkg/repository"
)

// Service represents the business logic layer, providing methods for handling user-related operations.
type Service struct {
	repository *repository.Repository
}

// NewService creates a new instance of Service with the given repository.
func NewService(repo *repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}

// AddUser adds a new user after validating the user's fields.
// It takes a pointer to a User model as an argument and returns the added User and any error encountered.
func (service Service) AddUser(user *model.User) (*model.User, error) {
	if user.Name == "" || user.Surname == "" || user.Age <= 0 {
		return nil, errors.New("user's fields can't be empty")
	}
	user, err := service.repository.AddUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user after validating the new user's fields.
// It takes a user ID and a pointer to a UserUpdateRequest model as arguments and returns the updated User and any error encountered.
func (service Service) UpdateUser(userID uint, newUser *model.UserUpdateRequest) (*model.User, error) {
	if newUser.Name == "" || newUser.Surname == "" || newUser.Age <= 0 {
		return nil, errors.New("user's fields can't be empty")
	}
	updatedUser, err := service.repository.UpdateUser(userID, newUser)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// DeleteUser deletes a user by user ID after validating the ID.
// It takes a user ID as an argument and returns any error encountered during the deletion.
func (service Service) DeleteUser(userID uint) error {
	if userID <= 0 {
		return errors.New("user ID can't be equal to or less than zero")
	}
	err := service.repository.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}

// GetUsers retrieves a list of all users from the repository.
// It returns a pointer to a slice of User models and any error encountered during the retrieval.
func (service Service) GetUsers() (*[]model.User, error) {
	users, err := service.repository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retrieves a user by user ID after validating the ID.
// It takes a user ID as an argument and returns a pointer to the User model and any error encountered during the retrieval.
func (service Service) GetUserByID(userID uint) (*model.User, error) {
	if userID <= 0 {
		return nil, errors.New("user ID can't be equal to or less than zero")
	}
	user, err := service.repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

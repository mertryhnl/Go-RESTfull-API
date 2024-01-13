package model

import "gorm.io/gorm"

// User represents the model for a user in the system.
// It includes fields like Name, Surname, and Age, along with gorm.Model for database operations.
type User struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

// UserUpdateRequest is a model representing the request structure for updating a user.
// It includes fields like Name, Surname, and Age that can be updated.
type UserUpdateRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

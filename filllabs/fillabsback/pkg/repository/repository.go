package repository

import (
	"fillabs_intern_project/pkg/model"
	"fmt"

	"gorm.io/gorm"
)

// Repository represents the data access layer, providing methods for interacting with the database.
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new instance of Repository with the given database connection.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// AddUser adds a new user to the database.
// It takes a pointer to a User model as an argument and returns the saved User and any error encountered.
func (repo *Repository) AddUser(user *model.User) (*model.User, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		fmt.Println("Error occurred when saving user: " + err.Error())
	}
	return user, nil
}

// UpdateUser updates an existing user in the database.
// It takes a user ID and a pointer to a UserUpdateRequest model as arguments and returns the updated User and any error encountered.
func (repo *Repository) UpdateUser(userID uint, newUser *model.UserUpdateRequest) (*model.User, error) {
	err := repo.db.Where("id = ?", userID).Find(&model.User{}).Error
	if err != nil {
		return nil, err
	}

	if err := repo.db.Model(&model.User{}).Where("id = ?", userID).Updates(&newUser).Error; err != nil {
		fmt.Println("Error occurred when saving user: " + err.Error())
		return nil, err
	}

	var user *model.User

	if err := repo.db.Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user from the database.
// It takes a user ID as an argument and returns any error encountered during the deletion.
func (repo *Repository) DeleteUser(userID uint) error {
	if err := repo.db.Unscoped().Where("id = ?", userID).Delete(&model.User{}).Error; err != nil {
		fmt.Println("Error occurred when deleting user: " + err.Error())
	}
	return nil
}

// GetUsers retrieves a list of all users from the database.
// It returns a pointer to a slice of User models and any error encountered during the retrieval.
func (repo *Repository) GetUsers() (*[]model.User, error) {
	var users *[]model.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retrieves a user from the database by its ID.
// It takes a user ID as an argument and returns a pointer to the User model and any error encountered during the retrieval.
func (repo *Repository) GetUserByID(userID uint) (*model.User, error) {
	var user *model.User
	if err := repo.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

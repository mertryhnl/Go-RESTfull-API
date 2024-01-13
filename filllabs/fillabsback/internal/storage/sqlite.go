package storage

import (
	"fillabs_intern_project/pkg/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// OpenDB opens a connection to the SQLite database and performs automatic migration
// for the User model.
// It returns a pointer to the gorm.DB instance and any error encountered during the process.
func OpenDB() (*gorm.DB, error) {
	// Open a connection to the SQLite database located at "user.db"
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		// Return an error if there is an issue opening the database connection
		return nil, err
	}

	// Automatically migrate the User model to create the corresponding table in the database
	db.AutoMigrate(&model.User{})

	// Return the gorm.DB instance and nil error upon successful database setup
	return db, nil
}

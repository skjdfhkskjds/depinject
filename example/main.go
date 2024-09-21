package main

import (
	"fmt"

	"github.com/skjdfhkskjds/depinject"
)

// UserService represents a service for managing users
type UserService struct {
	// Add any necessary fields
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{}
}

// DatabaseConnection represents a database connection
type DatabaseConnection struct {
	// Add any necessary fields
}

// NewDatabaseConnection creates a new DatabaseConnection
func NewDatabaseConnection() *DatabaseConnection {
	return &DatabaseConnection{}
}

// Application represents the main application
type Application struct {
	UserService *UserService
	DB          *DatabaseConnection
}

// NewApplication creates a new Application
func NewApplication(userService *UserService, db *DatabaseConnection) *Application {
	return &Application{
		UserService: userService,
		DB:          db,
	}
}

func main() {
	container := depinject.NewContainer()

	if err := container.Provide(
		NewDatabaseConnection,
		NewUserService,
		NewApplication,
	); err != nil {
		panic(err)
	}

	var app *Application
	if err := container.Invoke(&app); err != nil {
		panic(err)
	}

	fmt.Println("Application initialized successfully!")
}

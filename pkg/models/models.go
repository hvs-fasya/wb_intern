package models

// User user struct as per storage
type User struct {
	ID   int
	Name string
}

// CreateUserInput input struct for CreateUser
type CreateUserInput struct {
	Name string
}

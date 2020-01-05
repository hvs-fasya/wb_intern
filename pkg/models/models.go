package models

// package constants
const EmptyUserName = ""

// User user struct as per memstore
type User struct {
	ID   int
	Name string
}

// CreateUserInput input struct for CreateUser
type CreateUserInput struct {
	Name string
}

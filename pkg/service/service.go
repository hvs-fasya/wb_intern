package service

import (
	"github.com/hvs-fasya/wb_intern/pkg/models"
)

// UserService common user service interface
type UserService interface {
	CreateUser(input models.CreateUserInput) error
}

type notifier interface {
	NotifyCreateUser([]interface{})
}

type storage interface {
	NewUser(string) error
}

type userService struct {
	storage  storage
	notifier notifier
}

// CreateUser create user method, implements common user service interface
func (s *userService) CreateUser(inp models.CreateUserInput) error {
	if err := s.storage.NewUser(inp.Name); err != nil {
		return err
	}
	if inp.Type == "" {
		inp.Type = models.UserTypeSimple
	}
	//hidden behind facade action
	s.notifier.NotifyCreateUser([]interface{}{inp.Name})

	return nil
}

// NewUserService user service constructor
func NewUserService(n notifier, s storage) UserService {
	return &userService{
		storage:  s,
		notifier: n,
	}
}

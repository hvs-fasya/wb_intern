package service

import (
	"errors"
	"sync"

	"github.com/hvs-fasya/wb_intern/pkg/models"
)

// UserService common user service interface
type UserService interface {
	CreateUser(input models.CreateUserInput) error
}

type notifier interface {
	NotifyCreateUser([]string)
}

type userService struct {
	sync.Mutex
	storage  storage
	notifier notifier
}

type storage []models.User

// CreateUser create user method, implements common user service interface
func (s *userService) CreateUser(inp models.CreateUserInput) error {
	var curID int
	if inp.Name == "" {
		return errors.New("user name required")
	}
	s.Lock()
	for _, u := range s.storage {
		if curID < u.ID {
			curID = u.ID
		}
	}
	s.storage = append(s.storage, models.User{ID: curID + 1, Name: inp.Name})
	s.Unlock()
	//hidden behind facade action
	s.notifier.NotifyCreateUser([]string{inp.Name})

	return nil
}

// NewUserService user service constructor
func NewUserService(n notifier) UserService {
	s := make([]models.User, 0)
	return &userService{
		storage:  s,
		notifier: n,
	}
}

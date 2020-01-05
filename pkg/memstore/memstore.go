package memstore

import (
	"sync"

	"github.com/hvs-fasya/wb_intern/pkg/models"
)

// Storage common storage integace
type Storage interface {
	NewUser(string) error
}

type memstore struct {
	sync.Mutex
	users     []models.User
	curUserID int
}

// NewUser common Storage interface NewUser method implementation
func (ms *memstore) NewUser(userName string) error {
	if userName == models.EmptyUserName {
		return models.ErrUserNameRequired
	}
	ms.Lock()
	ms.users = append(ms.users, models.User{ID: ms.curUserID + 1, Name: userName})
	ms.curUserID++
	ms.Unlock()
	return nil
}

// NewMemStore memstore constructor
func NewMemStore() Storage {
	users := make([]models.User, 0)
	return &memstore{users: users}
}

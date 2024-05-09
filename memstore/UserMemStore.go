package memstore

import (
	"embedded-nextjs-poc/models"
	"errors"
)

var (
	NotFoundErr = errors.New("not found")
)

type UserMemStore struct {
	list map[string]models.User
}

func NewUserMemStore() *UserMemStore {
	list := make(map[string]models.User)
	return &UserMemStore{
		list,
	}
}

func (m UserMemStore) Add(User models.User) models.User {
	id := User.ID
	if id == "" {
		newID := models.NewUserId()
		User.ID = string(newID)
	}

	m.list[User.ID] = User
	return User
}

func (m UserMemStore) Get(id string) (models.User, error) {

	if val, ok := m.list[id]; ok {
		return val, nil
	}

	return models.User{}, NotFoundErr
}

func (m UserMemStore) List() (map[string]models.User, error) {
	return m.list, nil
}

func (m UserMemStore) Update(id string, User models.User) (models.User, error) {
	if _, ok := m.list[id]; ok {
		User.ID = id
		m.list[id] = User
		return User, nil
	}

	return models.User{}, NotFoundErr
}

func (m UserMemStore) Remove(id string) error {
	if _, ok := m.list[id]; ok {
		delete(m.list, id)
		return nil
	}
	return NotFoundErr

}

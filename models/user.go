package models

import "github.com/google/uuid"

type UserID string

func NewUserId() UserID {
	return UserID(uuid.New().String())
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

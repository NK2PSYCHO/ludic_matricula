package user

import "github.com/google/uuid"

type UserRepository interface {
	AddUser(user User) error
	UpdateUser(id uuid.UUID, user User) error
	DeleteUser(id uuid.UUID) error
	GetUser(id uuid.UUID) (User, error)
	GetUsers() ([]User, error)
}

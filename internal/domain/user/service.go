package user

import (
	"github.com/google/uuid"
)

var _ UserService = (*userService)(nil)

type UserService interface {
	AddUser(user User) error
	UpdateUser(id uuid.UUID, user User) error
	DeleteUser(id uuid.UUID) error
	GetUser(id uuid.UUID) (User, error)
	GetUsers(id uuid.UUID) ([]User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) AddUser(user User) error {
	panic("unimplemented")
}

func (u *userService) UpdateUser(id uuid.UUID, user User) error {
	panic("unimplemented")
}

func (u *userService) DeleteUser(id uuid.UUID) error {
	panic("unimplemented")
}

func (u *userService) GetUser(id uuid.UUID) (User, error) {
	panic("unimplemented")
}

func (u *userService) GetUsers(id uuid.UUID) ([]User, error) {
	panic("unimplemented")
}

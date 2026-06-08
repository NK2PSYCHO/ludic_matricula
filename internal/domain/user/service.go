package user

import (
	"github.com/google/uuid"
)

var _ UserService = (*userService)(nil)

type UserService interface {
	AddUser(req CreateUserReq) error
	UpdateUser(id uuid.UUID, user User) error
	DeleteUser(id uuid.UUID) error
	GetUser(id uuid.UUID) (UserDTO, error)
	GetUsers() ([]UserDTO, error)
	Login(req UserLoginReq) (UserDTO, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) AddUser(req CreateUserReq) error {
	panic("unimplemented")
}

func (u *userService) UpdateUser(id uuid.UUID, user User) error {
	panic("unimplemented")
}

func (u *userService) DeleteUser(id uuid.UUID) error {
	panic("unimplemented")
}

func (u *userService) GetUser(id uuid.UUID) (UserDTO, error) {
	panic("unimplemented")
}

func (u *userService) GetUsers() ([]UserDTO, error) {
	panic("unimplemented")
}

func (u *userService) Login(req UserLoginReq) (UserDTO, error) {
	panic("unimplemented")
}

package persistence

import (
	"ludic-matricula/internal/domain/user"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ user.UserRepository = (*userRepo)(nil)

type userRepo struct {
	db *bun.DB
}

func NewUserRepo(db *bun.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) AddUser(user user.User) error {
	panic("unimplemented")
}

func (u *userRepo) UpdateUser(id uuid.UUID, user user.User) error {
	panic("unimplemented")
}

func (u *userRepo) DeleteUser(id uuid.UUID) error {
	panic("unimplemented")
}

func (u *userRepo) GetUser(id uuid.UUID) (user.User, error) {
	panic("unimplemented")
}

func (u *userRepo) GetUsers() ([]user.User, error) {
	panic("unimplemented")
}

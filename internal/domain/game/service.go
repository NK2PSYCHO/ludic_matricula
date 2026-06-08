package game

import (
	"github.com/google/uuid"
)

var _ GameService = (*gameService)(nil)

type GameService interface {
	AddGame(user Game) error
	UpdateGame(id uuid.UUID, Game Game) error
	DeleteGame(id uuid.UUID) error
	GetGame(id uuid.UUID) (Game, error)
	GetGames() ([]Game, error)
}

type gameService struct {
	repo GameRepository
}

func NewGameService(repo GameRepository) *gameService {
	return &gameService{
		repo: repo,
	}
}

func (g *gameService) AddGame(user Game) error {
	panic("unimplemented")
}

func (g *gameService) UpdateGame(id uuid.UUID, Game Game) error {
	panic("unimplemented")
}

func (g *gameService) DeleteGame(id uuid.UUID) error {
	panic("unimplemented")
}

func (g *gameService) GetGame(id uuid.UUID) (Game, error) {
	panic("unimplemented")
}

func (g *gameService) GetGames() ([]Game, error) {
	panic("unimplemented")
}

package persistence

import (
	"ludic-matricula/internal/domain/game"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ game.GameRepository = (*gameRepo)(nil)

type gameRepo struct {
	db *bun.DB
}

func NewGameRepo(db *bun.DB) *gameRepo {
	return &gameRepo{
		db: db,
	}
}

func (g *gameRepo) AddGame(game game.Game) error {
	panic("unimplemented")
}

func (g *gameRepo) UpdateGame(id uuid.UUID, game game.Game) error {
	panic("unimplemented")
}

func (g *gameRepo) DeleteGame(id uuid.UUID) error {
	panic("unimplemented")
}

func (g *gameRepo) GetGame(id uuid.UUID) (game.Game, error) {
	panic("unimplemented")
}

func (g *gameRepo) GetGames() ([]game.Game, error) {
	panic("unimplemented")
}

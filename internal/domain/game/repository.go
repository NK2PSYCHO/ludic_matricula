package game

import "github.com/google/uuid"

type GameRepository interface {
	AddGame(game Game) error
	UpdateGame(id uuid.UUID, game Game) error
	DeleteGame(id uuid.UUID) error
	GetGame(id uuid.UUID) (Game, error)
	GetGames(id uuid.UUID) ([]Game, error)
}

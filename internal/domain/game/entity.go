package game

import (
	"github.com/google/uuid"
)

type Game struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Publisher string    `json:"publisher"`
	Year      string    `json:"year"`
	Platform  string    `json:"platform"`
}

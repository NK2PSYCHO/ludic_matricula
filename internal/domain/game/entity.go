package game

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Publisher string    `json:"publisher"`
	Year      string    `json:"year"`
	Platform  string    `json:"platform"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

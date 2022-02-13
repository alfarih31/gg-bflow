package meta_dto

import (
	"time"
)

type Item struct {
	ID        string    `json:"id"`
	Key       string    `json:"key"`
	Metadata  string    `json:"metadata"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

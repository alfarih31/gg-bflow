package meta_dto

import (
	"github.com/alfarih31/nb-go-keyvalue"
	"time"
)

type Item struct {
	ID        string            `json:"id"`
	Key       string            `json:"key"`
	Metadata  keyvalue.KeyValue `json:"metadata"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

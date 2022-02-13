package meta

import (
	"context"
	meta_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
)

type (
	Writer interface {
		Write(ctx context.Context, i meta_dto.Item) error
		Update(ctx context.Context, k string, i meta_dto.Item) error
	}

	Reader interface {
		Read(ctx context.Context, k string) (*meta_dto.Item, error)
	}

	Repo interface {
		Writer
		Reader
	}

	Doer struct {
		Writer
		Reader
	}
)

var Do Repo

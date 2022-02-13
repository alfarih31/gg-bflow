package buffer

import (
	"context"
	buffer_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
)

type (
	Writer interface {
		Write(ctx context.Context, key string, data []byte, force ...bool) error
		WriteNew(ctx context.Context, key string, data []byte) error
	}

	Reader interface {
		Read(ctx context.Context, key string) (s *buffer_dto.Stat, err error)
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

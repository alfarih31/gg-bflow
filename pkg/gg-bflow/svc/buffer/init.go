package buffer

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repo/buffer"
)

type BufferSvc interface {
	Read(ctx context.Context, key string) (*buffer_dto.WriteArg, error)
	Write(ctx context.Context, arg buffer_dto.WriteArg) error
}

type svc struct {
	repo buffer.BufferRepo
}

var Svc BufferSvc = &svc{
	repo: buffer.Repo,
}

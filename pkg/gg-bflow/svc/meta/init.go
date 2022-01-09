package meta

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repo/meta"
)

type MetaSvc interface {
	Read(ctx context.Context, key string) (*meta_dto.Item, error)
	Write(ctx context.Context, arg meta_dto.WriteArg) error
}

type svc struct {
	repo meta.MetaRepo
}

var Svc MetaSvc = &svc{
	repo: meta.Instance,
}

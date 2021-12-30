package buffer_svc

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/buffer/repo"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
)

func Read(ctx context.Context, key string) (*buffer_dto.WriteArg, error) {
	d, err := buffer_repo.Read(ctx, key)

	if err != nil {
		return nil, err
	}

	return &buffer_dto.WriteArg{
		Key:  d.Key,
		Data: d.Data,
	}, nil
}

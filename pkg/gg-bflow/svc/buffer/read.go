package buffer

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
)

func (s *svc) Read(ctx context.Context, key string) (*buffer_dto.WriteArg, error) {
	d, err := s.repo.Read(ctx, key)

	if err != nil {
		return nil, err
	}

	return &buffer_dto.WriteArg{
		Key:  d.Key,
		Data: d.Data,
	}, nil
}

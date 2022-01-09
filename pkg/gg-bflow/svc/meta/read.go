package meta

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
)

func (s *svc) Read(ctx context.Context, key string) (*meta_dto.Item, error) {
	i, err := s.repo.FindOne(ctx, key)
	if err != nil {
		return nil, err
	}

	return i, nil
}

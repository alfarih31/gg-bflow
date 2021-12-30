package meta_svc

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/meta/repo"
)

func Read(ctx context.Context, key string) (*meta_dto.Item, error) {
	i, err := meta_repo.FindOne(ctx, key)
	if err != nil {
		return nil, err
	}

	return i, nil
}

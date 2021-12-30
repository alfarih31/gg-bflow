package buffer_svc

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/buffer/repo"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
)

func Write(ctx context.Context, arg buffer_dto.WriteArg) (*buffer_dto.Stat, error) {
	stat, err := buffer_repo.Write(ctx, arg.Key, arg.Data, true)
	if err != nil {
		return nil, err
	}

	return stat, nil
}

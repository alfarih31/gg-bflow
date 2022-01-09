package buffer

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
)

func (s *svc) Write(ctx context.Context, arg buffer_dto.WriteArg) error {
	err := s.repo.Write(ctx, arg.Key, arg.Data, true)
	if err != nil {
		return err
	}

	return nil
}

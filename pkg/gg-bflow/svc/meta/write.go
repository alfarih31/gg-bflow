package meta

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/utils/datetime"
)

func (s *svc) Write(ctx context.Context, arg meta_dto.WriteArg) error {
	var i *meta_dto.Item

	// Get exist
	i, err := s.Read(ctx, arg.Key)
	if err != nil {
		return err
	}

	now := datetime.Now()
	if i == nil {
		i = &meta_dto.Item{
			Key:       arg.Key,
			Metadata:  arg.Meta,
			CreatedAt: now,
			UpdatedAt: now,
		}

		err = s.repo.Insert(ctx, i)
	} else {
		i.Metadata = arg.Meta
		i.UpdatedAt = now

		err = s.repo.UpdateOne(ctx, arg.Key, i)
	}

	return err
}

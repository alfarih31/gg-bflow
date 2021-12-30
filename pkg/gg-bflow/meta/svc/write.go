package meta_svc

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/meta/repo"
	"github.com/alfarih31/gg-bflow/pkg/utils/datetime"
)

func Write(ctx context.Context, arg meta_dto.WriteArg) error {
	var i *meta_dto.Item

	// Get exist
	i, err := Read(ctx, arg.Key)
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

		err = meta_repo.Insert(ctx, i)
	} else {
		i.Metadata = arg.Meta
		i.UpdatedAt = now

		err = meta_repo.Update(ctx, arg.Key, i)
	}

	return err
}

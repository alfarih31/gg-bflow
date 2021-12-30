package buffer_repo

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/memcache"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
	mc "github.com/bradfitz/gomemcache/memcache"
)

func Read(ctx context.Context, key string) (s *buffer_dto.Stat, err error) {
	i, err := memcache.Get(key)
	if err == mc.ErrCacheMiss {
		return &buffer_dto.Stat{
			Key: key,
		}, nil
	}

	if err != nil {
		return nil, err
	}

	return &buffer_dto.Stat{
		Key:  key,
		Data: i.Value,
		Exp:  int64(i.Expiration),
	}, nil
}

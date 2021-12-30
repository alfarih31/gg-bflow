package buffer_repo

import (
	"context"
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/memcache"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
	mc "github.com/bradfitz/gomemcache/memcache"
)

// Write will write data to memcached if it already exists. Use `force` to force write
func Write(ctx context.Context, key string, data []byte, force ...bool) (s *buffer_dto.Stat, err error) {
	f := false
	if len(force) > 0 {
		f = force[0]
	}

	i, err := memcache.Get(key)

	if err != nil && !f {
		return nil, err
	}

	if i == nil && f {
		return WriteNew(ctx, key, data)
	}

	i.Value = data
	i.Expiration = int32(configs.GGBFlow.BufferExp)

	err = memcache.Replace(i)
	if err != nil {
		return nil, err
	}

	return &buffer_dto.Stat{
		Key:  key,
		Data: data,
		Exp:  int64(i.Expiration),
	}, nil
}

func WriteNew(ctx context.Context, key string, data []byte) (s *buffer_dto.Stat, err error) {
	i := &mc.Item{
		Key:        key,
		Value:      data,
		Expiration: int32(configs.GGBFlow.BufferExp),
	}

	err = memcache.Add(i)
	if err != nil {
		return nil, err
	}

	return &buffer_dto.Stat{
		Key:  key,
		Data: data,
		Exp:  int64(i.Expiration),
	}, nil
}

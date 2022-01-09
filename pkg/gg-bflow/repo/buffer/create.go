package buffer

import (
	"context"
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/memcache"
	mc "github.com/bradfitz/gomemcache/memcache"
)

type Create interface {
	Write(ctx context.Context, key string, data []byte, force ...bool) error
	WriteNew(ctx context.Context, key string, data []byte) error
}

type create struct {
}

// Write will write data to memcached if it already exists. Use `force` to force write
func (c *create) Write(ctx context.Context, key string, data []byte, force ...bool) (err error) {
	f := false
	if len(force) > 0 {
		f = force[0]
	}

	i, err := memcache.Get(key)

	if err != nil && !f {
		return err
	}

	if i == nil && f {
		return c.WriteNew(ctx, key, data)
	}

	i.Value = data
	i.Expiration = int32(configs.GGBFlow.BufferExp)

	err = memcache.Replace(i)
	if err != nil {
		return err
	}

	return nil
}

func (c *create) WriteNew(ctx context.Context, key string, data []byte) (err error) {
	i := &mc.Item{
		Key:        key,
		Value:      data,
		Expiration: int32(configs.GGBFlow.BufferExp),
	}

	err = memcache.Set(i)
	if err != nil {
		return err
	}

	return nil
}

func NewCreate() Create {
	return new(create)
}

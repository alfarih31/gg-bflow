package bufferImpl

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/memcache"
	buffer_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/buffer"
	mc "github.com/bradfitz/gomemcache/memcache"
)

var _ buffer.Reader = new(reader)

type reader struct {
}

func NewReader() buffer.Reader {
	return new(reader)
}

func (r *reader) Read(ctx context.Context, key string) (s *buffer_dto.Stat, err error) {
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

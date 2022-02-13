package metaMock

import (
	"context"
	meta_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta"
)

func NewReaderMock() meta.Reader {
	return new(reader)
}

type reader struct {
}

func (r reader) Read(ctx context.Context, k string) (*meta_dto.Item, error) {
	return &meta_dto.Item{}, nil
}

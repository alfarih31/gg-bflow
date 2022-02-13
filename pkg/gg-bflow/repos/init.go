package repos

import (
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/buffer"
	bufferImpl "github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/buffer/impl"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta"
	metaImpl "github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta/impl"
	metaMock "github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta/mocks"
)

func Init() error {
	buffer.Do = &buffer.Doer{
		Writer: bufferImpl.NewWriter(),
		Reader: bufferImpl.NewReader(),
	}

	meta.Do = &meta.Doer{
		Writer: metaImpl.NewWriter(),
		Reader: metaMock.NewReaderMock(),
	}

	return nil
}

package repos

import (
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/buffer"
	bufferImpl "github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/buffer/impl"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta"
	metaImpl "github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta/impl"
)

func Init() error {
	buffer.Do = &buffer.Doer{
		Writer: bufferImpl.NewWriter(),
		Reader: bufferImpl.NewReader(),
	}

	meta.Do = &meta.Doer{
		Writer: metaImpl.NewWriter(),
		Reader: metaImpl.NewReader(),
	}

	return nil
}

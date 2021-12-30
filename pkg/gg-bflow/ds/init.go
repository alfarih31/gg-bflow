package ds

import (
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/memcache"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo"
)

func Init() error {
	if err := mongo.Init(); err != nil {
		return err
	}

	if err := memcache.Init(); err != nil {
		return err
	}

	return nil
}

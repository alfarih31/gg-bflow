package memcache

import (
	"fmt"
	"github.com/alfarih31/gg-bflow/configs"
	mc "github.com/bradfitz/gomemcache/memcache"
)

var c *mc.Client

func Replace(i *mc.Item) error {
	return c.Replace(i)
}

func Add(i *mc.Item) error {
	return c.Add(i)
}

func Get(k string) (*mc.Item, error) {
	return c.Get(k)
}

func Init() error {
	cfg := configs.Memcache

	c = mc.New(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))

	return nil
}

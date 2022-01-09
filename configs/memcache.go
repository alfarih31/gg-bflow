package configs

import (
	_env "github.com/alfarih31/nb-go-env"
	"time"
)

type memcache struct {
	Host          string        `validate:"required,min=0"`
	Port          int           `validate:"required,gte=0"`
	Timeout       time.Duration `validate:"required,gte=0"`
	MaxConcurrent int           `validate:"required,gte=0"`
}

var Memcache = new(memcache)

func (m *memcache) Load(env _env.Env) error {
	if v, err := env.GetString("MEMCACHE_HOST"); err != nil {
		return err
	} else {
		m.Host = v
	}

	if v, err := env.GetInt("MEMCACHE_PORT"); err != nil {
		return err
	} else {
		m.Port = v
	}

	if v, err := env.GetInt("MEMCACHE_TIMEOUT", 100); err != nil {
		return err
	} else {
		m.Timeout = time.Millisecond * time.Duration(v)
	}

	if v, err := env.GetInt("MEMCACHE_MAX_CONCURRENT", 2); err != nil {
		return err
	} else {
		m.MaxConcurrent = v
	}

	return nil
}

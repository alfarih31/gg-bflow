package configs

import _env "github.com/alfarih31/nb-go-env"

type memcache struct {
	Host string
	Port int
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

	return nil
}

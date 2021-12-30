package configs

import (
	_env "github.com/alfarih31/nb-go-env"
)

var Env _env.Env

func Init(envFile string) error {
	env, err := _env.LoadEnv(envFile, true)
	if err != nil {
		return err
	}

	Env = env

	if err := Memcache.Load(env); err != nil {
		return err
	}

	if err := GGBFlow.Load(env); err != nil {
		return err
	}

	if err := Mongo.Load(env); err != nil {
		return err
	}

	return nil
}

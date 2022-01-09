package configs

import _env "github.com/alfarih31/nb-go-env"

type mongo struct {
	Database   string
	Host       string
	Port       int
	User       string
	Pass       string
	AuthSource string
}

var Mongo = new(mongo)

func (m *mongo) Load(env _env.Env) error {
	if v, e := env.GetString("MONGO_DATABASE"); e != nil {
		return e
	} else {
		m.Database = v
	}

	if v, e := env.GetString("MONGO_HOST", "localhost"); e != nil {
		return e
	} else {
		m.Host = v
	}

	if v, e := env.GetInt("MONGO_PORT", 27017); e != nil {
		return e
	} else {
		m.Port = v
	}

	if v, e := env.GetString("MONGO_USER"); e != nil {
		return e
	} else {
		m.User = v
	}

	if v, e := env.GetString("MONGO_PASS"); e != nil {
		return e
	} else {
		m.Pass = v
	}

	if v, e := env.GetString("MONGO_AUTHSOURCE"); e != nil {
		return e
	} else {
		m.AuthSource = v
	}

	return nil
}

package configs

import (
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/validator"
	_env "github.com/alfarih31/nb-go-env"
)

type ggbflow struct {
	Port             int      `validate:"required,gt=0"`
	APIKey           string   `validate:"required,min=0"`
	AuthorizedClient []string `validate:"required"`
	ServerName       string
	BufferSizeLimit  int `validate:"required,gt=0"`
	BufferExp        int `validate:"required,gte=0"`
}

var GGBFlow = new(ggbflow)

func (b *ggbflow) Load(env _env.Env) error {
	if grpcPort, err := env.GetInt("BFLOW_PORT", 50051); err != nil {
		return err
	} else {
		b.Port = grpcPort
	}

	if grpcKey, err := env.GetString("BFLOW_API_KEY"); err != nil {
		return err
	} else {
		b.APIKey = grpcKey
	}

	if v, err := env.GetStringArr("BFLOW_AUTHORIZED_CLIENT"); err != nil {
		return err
	} else {
		b.AuthorizedClient = v
	}

	if v, err := env.GetString("BFLOW_HOSTNAME", ""); err != nil {
		return err
	} else {
		b.ServerName = v
	}

	if v, err := env.GetInt("BFLOW_BUFFER_SIZE_LIMIT"); err != nil {
		return err
	} else {
		b.BufferSizeLimit = v
	}

	if v, err := env.GetInt("BFLOW_BUFFER_EXP"); err != nil {
		return err
	} else {
		b.BufferExp = v
	}

	return validator.Validate(b)
}

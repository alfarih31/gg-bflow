package ctrl

import (
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc"
)

func Start() error {
	return grpc_ctrl.Start()
}

func Init() error {
	if err := grpc_ctrl.Init(); err != nil {
		return err
	}

	return nil
}

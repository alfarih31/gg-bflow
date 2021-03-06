package gg_bflow

import (
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/logger"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos"
)

func Init(envFile ...string) error {
	ef := ".env"
	if len(envFile) > 0 {
		ef = envFile[0]
	}

	if err := configs.Init(ef); err != nil {
		return err
	}

	if err := ds.Init(); err != nil {
		return err
	}

	if err := repos.Init(); err != nil {
		return err
	}

	if err := ctrl.Init(); err != nil {
		return err
	}

	return nil
}

func Start() {
	if err := ctrl.Start(); err != nil {
		logger.Log.Error(err)
		return
	}
}

package gg_bflow

import (
	app "github.com/alfarih31/gg-bflow"
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/logger"
)

func Init(envFile ...string) error {
	ef := ".env"
	if len(envFile) > 0 {
		ef = envFile[0]
	}

	if err := configs.Init(ef); err != nil {
		return err
	}

	if v, err := configs.Env.GetString("APP_VERSION", app.Meta.AppVersion); err != nil {
		return err
	} else {
		app.Meta.AppVersion = v
	}

	if err := ds.Init(); err != nil {
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

package mongo

import (
	"context"
	"fmt"
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo/models"
	"github.com/qiniu/qmgo"
)

var (
	Session *qmgo.Session
	Query   *models.Query
)

func Init() error {
	cfg := configs.Mongo
	client, err := qmgo.NewClient(context.TODO(), &qmgo.Config{
		Uri: fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port),
		Auth: &qmgo.Credential{
			Username:   cfg.User,
			Password:   cfg.Pass,
			AuthSource: cfg.AuthSource,
		},
	})

	if err != nil {
		return err
	}

	Session, err = client.Session()
	if err != nil {
		return err
	}

	db := client.Database(cfg.Database)
	Query = models.Use(db)

	return nil
}

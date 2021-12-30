package models

import (
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo/model"
	"github.com/qiniu/qmgo"
)

type Query struct {
	Meta *qmgo.Collection
}

func Use(d *qmgo.Database) *Query {
	return &Query{
		Meta: d.Collection(model.CollNameMeta),
	}
}

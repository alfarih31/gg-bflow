package metaImpl

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo/model"
	meta_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type writer struct {
}

func NewWriter() meta.Writer {
	return new(writer)
}

func (w *writer) Write(ctx context.Context, i meta_dto.Item) error {
	item := &model.Meta{
		Key:       i.Key,
		Metadata:  i.Metadata,
		CreatedAt: primitive.NewDateTimeFromTime(i.CreatedAt),
		UpdatedAt: primitive.NewDateTimeFromTime(i.UpdatedAt),
	}

	_, err := mongo.Query.Meta.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (w *writer) Update(ctx context.Context, k string, i meta_dto.Item) error {
	item := &model.Meta{
		Metadata:  i.Metadata,
		CreatedAt: primitive.NewDateTimeFromTime(i.CreatedAt),
		UpdatedAt: primitive.NewDateTimeFromTime(i.UpdatedAt),
	}

	err := mongo.Query.Meta.UpdateOne(ctx, bson.M{"key": k}, bson.M{"$set": item})
	if err != nil {
		return err
	}

	return nil
}

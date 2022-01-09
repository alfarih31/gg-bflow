package meta

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo/model"
	meta_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Update interface {
	UpdateOne(ctx context.Context, key string, i *meta_dto.Item) error
}

type update struct {
}

func (update) UpdateOne(ctx context.Context, key string, i *meta_dto.Item) error {
	item := &model.Meta{
		Metadata:  bson.M(i.Metadata),
		CreatedAt: primitive.NewDateTimeFromTime(i.CreatedAt),
		UpdatedAt: primitive.NewDateTimeFromTime(i.UpdatedAt),
	}

	err := mongo.Query.Meta.UpdateOne(ctx, bson.M{"key": key}, bson.M{"$set": item})
	if err != nil {
		return err
	}

	return nil
}

func NewUpdate() Update {
	return new(update)
}

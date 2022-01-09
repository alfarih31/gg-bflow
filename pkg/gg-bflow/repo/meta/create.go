package meta

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo/model"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Create interface {
	Insert(ctx context.Context, i *meta_dto.Item) error
}

type create struct {
}

func (create) Insert(ctx context.Context, i *meta_dto.Item) error {
	item := &model.Meta{
		Key:       i.Key,
		Metadata:  bson.M(i.Metadata),
		CreatedAt: primitive.NewDateTimeFromTime(i.CreatedAt),
		UpdatedAt: primitive.NewDateTimeFromTime(i.UpdatedAt),
	}

	_, err := mongo.Query.Meta.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func NewCreate() Create {
	return new(create)
}

package meta_repo

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo/model"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	keyvalue "github.com/alfarih31/nb-go-keyvalue"
	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
)

func FindOne(ctx context.Context, k string) (*meta_dto.Item, error) {
	i := new(model.Meta)
	err := mongo.Query.Meta.Find(ctx, bson.M{"key": k}).One(i)

	if err == mg.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &meta_dto.Item{
		ID:        i.ID.String(),
		Key:       i.Key,
		Metadata:  keyvalue.KeyValue(i.Metadata),
		CreatedAt: i.CreatedAt.Time(),
		UpdatedAt: i.UpdatedAt.Time(),
	}, nil
}

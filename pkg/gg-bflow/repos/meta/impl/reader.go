package metaImpl

import (
	"context"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ds/mongo/model"
	meta_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta"
	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
)

type reader struct {
}

func NewReader() meta.Reader {
	return new(reader)
}

func (r *reader) Read(ctx context.Context, k string) (*meta_dto.Item, error) {
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
		Metadata:  i.Metadata,
		CreatedAt: i.CreatedAt.Time(),
		UpdatedAt: i.UpdatedAt.Time(),
	}, nil
}

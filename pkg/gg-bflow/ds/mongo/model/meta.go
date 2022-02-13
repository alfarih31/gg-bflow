package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollNameMeta = "meta"

type Meta struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Key       string             `bson:"key,omitempty" json:"key"`
	Metadata  string             `bson:"metadata,omitempty" json:"metadata"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}

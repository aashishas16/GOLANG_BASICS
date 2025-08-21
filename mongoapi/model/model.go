package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"` //bson tag for MongoDB, json tag for JSON serialization
	Movie   string             `bson:"movie" json:"movie"`
	Watched bool               `bson:"watched" json:"watched"`
}

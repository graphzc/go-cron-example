package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Message struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Body      string        `bson:"body" json:"body"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}

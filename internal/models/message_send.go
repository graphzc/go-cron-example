package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MessageSendStatus string

const (
	MessageSendStatusSuccess MessageSendStatus = "SUCCESS"
)

type MessageSend struct {
	ID        bson.ObjectID     `bson:"_id" json:"id"`
	MessageID bson.ObjectID     `bson:"message_id" json:"message_id"`
	Status    MessageSendStatus `bson:"status" json:"status"`
	CreatedAt time.Time         `bson:"created_at" json:"created_at"`
}

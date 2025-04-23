package repositories

import (
	"context"
	"time"

	"github.com/graphzc/go-cron-example/internal/config"
	"github.com/graphzc/go-cron-example/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MessageSendRepository interface {
	Create(ctx context.Context, messageID bson.ObjectID, status models.MessageSendStatus) error
	ListAllSuccess(ctx context.Context) ([]*models.MessageSend, error)
}

type messageSendRepositoryImpl struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewMessageSendRepository(config *config.Config, mongoClient *mongo.Client) MessageSendRepository {
	collection := mongoClient.Database(config.Mongo.Database).Collection("message_sends")
	return &messageSendRepositoryImpl{
		mongoClient: mongoClient,
		collection:  collection,
	}
}

func (m *messageSendRepositoryImpl) Create(ctx context.Context, messageID bson.ObjectID, status models.MessageSendStatus) error {
	_, err := m.collection.InsertOne(ctx, models.MessageSend{
		ID:        bson.NewObjectID(),
		MessageID: messageID,
		Status:    status,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *messageSendRepositoryImpl) ListAllSuccess(ctx context.Context) ([]*models.MessageSend, error) {
	cursor, err := m.collection.Find(ctx, bson.M{"status": models.MessageSendStatusSuccess})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	messages := make([]*models.MessageSend, 0)
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

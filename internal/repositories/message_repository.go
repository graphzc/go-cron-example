package repositories

import (
	"context"

	"github.com/graphzc/go-cron-example/internal/config"
	"github.com/graphzc/go-cron-example/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MessageRepository interface {
	ListAll(ctx context.Context) ([]*models.Message, error)
	ListNotIn(ctx context.Context, messageIDs []bson.ObjectID) ([]*models.Message, error)
}

type messageRepositoryImpl struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewMessageRepository(config *config.Config, mongoClient *mongo.Client) MessageRepository {
	return &messageRepositoryImpl{
		mongoClient: mongoClient,
		collection:  mongoClient.Database(config.Mongo.Database).Collection("messages"),
	}
}

func (m *messageRepositoryImpl) ListAll(ctx context.Context) ([]*models.Message, error) {
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	messages := make([]*models.Message, 0)
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (m *messageRepositoryImpl) ListNotIn(ctx context.Context, messageIDs []bson.ObjectID) ([]*models.Message, error) {
	cursor, err := m.collection.Find(ctx, bson.M{
		"_id": bson.M{
			"$nin": messageIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	messages := make([]*models.Message, 0)
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

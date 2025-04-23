package database

import (
	"context"

	tnMongo "github.com/cnc-csku/task-nexus-go-lib/mongo"
	"github.com/graphzc/go-cron-example/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewMongoClient(ctx context.Context, configs *config.Config) *mongo.Client {
	return tnMongo.NewMongoClient(
		ctx,
		configs.Mongo.URI,
		configs.Mongo.Database,
	)
}

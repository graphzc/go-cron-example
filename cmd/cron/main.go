package main

import (
	"context"
	"flag"

	"github.com/graphzc/go-cron-example/cmd/cron/jobs"
	"github.com/graphzc/go-cron-example/internal/config"
	"github.com/graphzc/go-cron-example/internal/infrastructure/database"
	"github.com/graphzc/go-cron-example/internal/infrastructure/line"
	"github.com/graphzc/go-cron-example/internal/repositories"
	"github.com/graphzc/go-cron-example/internal/services"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.NewConfig()

	ctx := context.Background()

	mongoClient := database.NewMongoClient(ctx, config)

	lineClient := line.NewMockedLineClient("FAKE_ACCESS_TOKEN")

	messageRepository := repositories.NewMessageRepository(config, mongoClient)
	messageSendRepository := repositories.NewMessageSendRepository(config, mongoClient)

	messageService := services.NewMessageService(messageRepository, messageSendRepository, lineClient)

	boardCastCron := jobs.NewBoardcastMessageCron(messageService)

	_ = boardCastCron

	var cronName string

	flag.StringVar(&cronName, "name", "default", "")
	flag.Parse()

	runCtx := context.Background()

	if cronName == "line-board-cast" {
		boardCastCron.Run(runCtx)
	} else {
		logrus.Error("Invalid cron name")
	}
}

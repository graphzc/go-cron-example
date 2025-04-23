package jobs

import (
	"context"

	"github.com/graphzc/go-cron-example/internal/services"
	"github.com/sirupsen/logrus"
)

type BoardcastMessageCron struct {
	messageService services.MessageService
}

func NewBoardcastMessageCron(
	messageService services.MessageService,
) *BoardcastMessageCron {
	return &BoardcastMessageCron{
		messageService: messageService,
	}
}

func (b *BoardcastMessageCron) Run(ctx context.Context) {
	// Get all messages from the database
	unprocessedMessages, err := b.messageService.ListAllUnprocessedMessages(ctx)
	if err != nil {
		logrus.Errorf("Failed to list messages: %v", err)
		return
	}

	success := 0
	failed := 0

	logrus.Info("Start process messages")

	// Send each message to the LINE client
	for _, message := range unprocessedMessages {
		err := b.messageService.Boardcast(ctx, message)
		if err != nil {
			failed++
			logrus.Errorf("fail to boardcast message to line for message id %d", message.ID)
			continue
		}

		success++
	}

	logrus.Infof("Success: %d", success)
	logrus.Infof("Failed:  %d", failed)
}

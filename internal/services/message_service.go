package services

import (
	"context"

	"github.com/graphzc/go-cron-example/internal/infrastructure/line"
	"github.com/graphzc/go-cron-example/internal/models"
	"github.com/graphzc/go-cron-example/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MessageService interface {
	ListAllMessages(ctx context.Context) ([]*models.Message, error)
	Boardcast(ctx context.Context, message *models.Message) error
	ListAllUnprocessedMessages(ctx context.Context) ([]*models.Message, error)
}

type messageServiceImpl struct {
	messageRepository     repositories.MessageRepository
	messageSendRepository repositories.MessageSendRepository
	lineClient            line.LineClient
}

func NewMessageService(
	messageRepository repositories.MessageRepository,
	messageSendRepository repositories.MessageSendRepository,
	line line.LineClient,
) MessageService {
	return &messageServiceImpl{
		lineClient:            line,
		messageRepository:     messageRepository,
		messageSendRepository: messageSendRepository,
	}
}

func (m *messageServiceImpl) ListAllMessages(ctx context.Context) ([]*models.Message, error) {
	return m.messageRepository.ListAll(ctx)
}

func (m *messageServiceImpl) Boardcast(ctx context.Context, message *models.Message) error {
	err := m.lineClient.Boardcast(message.Body)
	if err != nil {
		return err
	}

	err = m.messageSendRepository.Create(ctx, message.ID, models.MessageSendStatusSuccess)
	if err != nil {
		return err
	}

	return nil
}

func (m *messageServiceImpl) ListAllUnprocessedMessages(ctx context.Context) ([]*models.Message, error) {
	successedMessages, err := m.messageSendRepository.ListAllSuccess(ctx)
	if err != nil {
		return nil, err
	}

	successedMessageIDs := make([]bson.ObjectID, len(successedMessages))
	for i, message := range successedMessages {
		successedMessageIDs[i] = message.MessageID
	}

	unprocessedMessages, err := m.messageRepository.ListNotIn(ctx, successedMessageIDs)
	if err != nil {
		return nil, err
	}

	return unprocessedMessages, nil
}

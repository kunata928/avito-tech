package tg

import (
	"context"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/logger"
	"go.uber.org/zap"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/model/messages"
)

type TokenGetter interface {
	TokenTg() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tokenGetter TokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.TokenTg())
	if err != nil {
		logger.Error("New bot API error", zap.Error(err))
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userID int64) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userID, text))
	if err != nil {
		logger.Error("Send message to client err", zap.Error(err))
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) ListenUpdates(ctx context.Context, msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates := c.client.GetUpdatesChan(u)
	logger.Info("Listening for messages...") //, zap.Int("port", 8080)

	go func() {
		<-ctx.Done()
		c.client.StopReceivingUpdates()
		logger.Info("Stop listening updates")
	}()

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := msgModel.IncomingMessage(ctx, messages.Message{
				Text:    update.Message.Text,
				UserID:  update.Message.From.ID,
				Command: update.Message.Command(),
			})
			if err != nil {
				logger.Error("Error processing message:", zap.Error(err))
			}
		}
	}
}

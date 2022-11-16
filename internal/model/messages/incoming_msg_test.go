package messages

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/metrics"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/logger"
	mocks "gitlab.ozon.dev/kunata928/telegramBot/internal/mocks/messages"
)

var testCases = []struct {
	name    string // название тест кейса
	msg     string // текст сообщения
	command string // команда
	answer  string // ответ
}{
	{
		name:    "help command",
		msg:     "/help",
		command: "help",
		answer:  helpMsg,
	},
	//{
	//	name:    "start command",
	//	msg:     "/start",
	//	command: "start",
	//	answer:  startMsg,
	//},
	{
		name:    "default answer",
		msg:     "/papapa",
		command: "",
		answer:  defaultMsg,
	},
	// ...
}

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	metrics.CommandStartTotal.Inc()
	logger.Info("test")
	ctx := context.Background()

	span, ctx := opentracing.StartSpanFromContext(ctx, defineCommand("test"))
	defer span.Finish()

	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	storage := mocks.NewMockStorage(ctrl)
	ratecl := mocks.NewMockRateClient(ctrl)
	cache := mocks.NewMockLRU(ctrl)
	model := New(sender, storage, ratecl, cache)
	for _, tc := range testCases {
		// тут запускаем каждый тест кейс по отдельности,
		// чтобы если упал какой либо тест кейс, знать какой именно
		t.Run(tc.name, func(t *testing.T) {
			//storage.EXPECT().InitClient(ctx, int64(123)) //.Return(nil).AnyTimes() //gomock.Any()
			sender.EXPECT().SendMessage(tc.answer, int64(123))
			err := model.IncomingMessage(ctx, Message{
				Text:    tc.msg,
				UserID:  123,
				Command: tc.command,
			})

			assert.NoError(t, err)
		})
	}
}

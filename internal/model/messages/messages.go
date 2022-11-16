package messages

import (
	"context"
	"go.uber.org/zap"
	"time"

	"github.com/shopspring/decimal"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/clients/rate"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/database"
)

var Logger *zap.Logger

type Message struct {
	Text    string
	UserID  int64
	Command string
}

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Storage interface {
	AddExpense(ctx context.Context, userID int64, expense *database.Expense) error
	GetClientExpenses(ctx context.Context, userID int64, fromDate time.Time) ([]*database.Expense, error)
	SetRate(ctx context.Context, rates *database.Rates, date time.Time) error
	GetRate(ctx context.Context, name string, date time.Time) (decimal.Decimal, error)
	RefreshClientCurrency(ctx context.Context, userID int64, currency string) error
	GetClientCurrency(ctx context.Context, userID int64) (string, error)
	InitClient(ctx context.Context, userID int64)
	RefreshClientLimit(ctx context.Context, userID int64, amount decimal.Decimal) error
}

type RateClient interface {
	GetRateDate(ctx context.Context, date time.Time) (*rate.Data, error)
}

type LRU interface {
	Add(key string, value interface{}) bool
	Get(key string) interface{}
	Remove(key string) bool
	Len() int
}

type Model struct {
	client     MessageSender
	storage    Storage
	rateClient RateClient
	cacheLRU   LRU
}

func New(client MessageSender, storage Storage, rateClient RateClient, cacheLRU LRU) *Model {
	return &Model{
		client:     client,
		storage:    storage,
		rateClient: rateClient,
		cacheLRU:   cacheLRU,
	}
}

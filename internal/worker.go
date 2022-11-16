package internal

import (
	"context"
	"time"
)

type CurrencyExchangeRateUpdater interface {
	UpdateCurrencyExchangeRatesOn(ctx context.Context, time time.Time) error
}

type ConfigGetter interface {
	GetFrequencyExchangeRate() time.Duration
}

type CurrencyExchangeRateWorker struct {
	updater CurrencyExchangeRateUpdater
	config  ConfigGetter
}

func NewCurrencyExchangeRateWorker(updater CurrencyExchangeRateUpdater, config ConfigGetter) *CurrencyExchangeRateWorker {
	return &CurrencyExchangeRateWorker{
		updater: updater,
		config:  config,
	}
}

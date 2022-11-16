package database

import (
	"strings"

	"github.com/shopspring/decimal"
)

type RateType string

const (
	USD RateType = "USD"
	RUB RateType = "RUB"
	EUR RateType = "EUR"
	CNY RateType = "CNY"
)

type Rates map[RateType]decimal.Decimal

func (r Rates) GetByCurrency(curr string) decimal.Decimal {
	if v, ok := r[RateType(strings.ToUpper(curr))]; ok {
		return v
	}
	return decimal.NewFromInt(0)
}

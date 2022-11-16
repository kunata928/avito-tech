package database

import (
	"github.com/shopspring/decimal"
)

type User struct {
	ID       int64
	Currency string
	Limit    decimal.Decimal
}

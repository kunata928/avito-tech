package database

import (
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct {
	Amount decimal.Decimal
	//Currency string
	Category string
	Date     time.Time
}

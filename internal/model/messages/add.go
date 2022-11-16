package messages

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/clients/rate"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/database"
)

var addRegex = regexp.MustCompile(`\s*(\d+|\d+.\d+)\s+([a-z0-9]+)\s+(\d{1,2}/\d{2}/\d{4})\s*`)

func dataToRates(data *rate.Data) *database.Rates {
	resRates := database.Rates{}
	resRates[database.RUB] = decimal.NewFromInt(1)
	resRates[database.USD] = decimal.NewFromInt(1).Div(data.Rates["RUB"])
	resRates[database.EUR] = resRates[database.USD].Mul(data.Rates["EUR"])
	resRates[database.CNY] = resRates[database.USD].Mul(data.Rates["CNY"])
	return &resRates
}

func (m *Model) getRatesAndSetToStorage(ctx context.Context, date time.Time) (*database.Rates, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getRatesAndSetToStorage")
	defer span.Finish()

	data, err := m.rateClient.GetRateDate(ctx, date)
	if err != nil {
		return nil, err
	}
	rates := dataToRates(data)
	if err := m.storage.SetRate(ctx, rates, date); err != nil {
		return nil, err
	}
	return rates, nil
}

func (m *Model) parsingCurrency(ctx context.Context, date time.Time, userID int64) (decimal.Decimal, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "parsingCurrency")
	defer span.Finish()

	curr, err := m.storage.GetClientCurrency(ctx, userID)
	if err != nil {
		return decimal.NewFromInt(0), errors.New(internalErrMsg)
	}
	rate, err := m.storage.GetRate(ctx, strings.ToUpper(curr), date)
	if err != nil {
		if !errors.Is(err, database.ErrNotFound) {
			return decimal.NewFromInt(0), errors.New(internalErrMsg)
		}
	}
	if !rate.IsZero() {
		return rate, nil
	}
	rates, err := m.getRatesAndSetToStorage(ctx, date)
	if err != nil {
		return decimal.NewFromInt(0), err
	}
	return rates.GetByCurrency(curr), nil
}

func (m *Model) parsingAdd(ctx context.Context, msg Message) (*database.Expense, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "parsingAdd")
	defer span.Finish()

	clientTxt := msg.Text[len("/add"):]
	if !addRegex.MatchString(clientTxt) {
		log.Println("Couldn't parse client request")
		return nil, errors.New("Couldn't parse client request")
	}
	groups := addRegex.FindStringSubmatch(clientTxt)
	date, err := time.Parse("2/1/2006", groups[3])
	if err != nil {
		log.Println("Couldn't parse date")
		return nil, errors.New("Couldn't parse date")
	}
	amount, err := decimal.NewFromString(groups[1])
	if err != nil {
		log.Println("Panic err")
		return nil, errors.New("Panic error")
	}
	rate, err := m.parsingCurrency(ctx, date, msg.UserID)
	if err != nil {
		return nil, err
	}

	return &database.Expense{
		Amount:   amount.Div(rate),
		Category: strings.ToLower(groups[2]),
		Date:     date,
	}, nil
}

func (m *Model) addExpense(ctx context.Context, msg Message) string {
	span, ctx := opentracing.StartSpanFromContext(ctx, "addExpense")
	defer span.Finish()

	log.Println(msg.Text[len("/add"):])
	expense, err := m.parsingAdd(ctx, msg)
	log.Println(expense)
	if err != nil {
		return parseErrMsg
	}
	if err := m.storage.AddExpense(ctx, msg.UserID, expense); err != nil {
		if err == database.ErrReachLimit {
			return "You reached your limit! Change /limit or less expense."
		}
		log.Println(err)
		return internalErrMsg
	}
	return "Successfully add your expense!"
}

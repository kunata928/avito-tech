package messages

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func getSubDate(period string) time.Time {
	dateNow := time.Now()
	subDate := time.Now()
	switch strings.ToLower(period) {
	case "week":
		subDate = dateNow.AddDate(0, 0, -8)
	case "month":
		subDate = dateNow.AddDate(0, -1, -1)
	case "year":
		subDate = dateNow.AddDate(-1, 0, -1)
	}
	return subDate
}

func (m *Model) processingReport(ctx context.Context, userID int64, period string) (string, error) {
	report := make(map[string]decimal.Decimal)
	curr, err := m.storage.GetClientCurrency(ctx, userID)
	resMsg := fmt.Sprintf("Your expenses on last %s in %s is:\n", period, curr)
	if err != nil {
		return "", errors.New(internalErrMsg)
	}
	expenses, err := m.storage.GetClientExpenses(ctx, userID, getSubDate(period))
	if err != nil {
		return "", err
	}

	for _, exp := range expenses {
		rate, err := m.storage.GetRate(ctx, curr, exp.Date)
		if err != nil {
			rates, err := m.getRatesAndSetToStorage(ctx, exp.Date)
			if err != nil {
				return "", err
			}
			rate = rates.GetByCurrency(curr)
		}
		report[exp.Category] = report[exp.Category].Add(exp.Amount.Mul(rate))
	}
	for cat, sum := range report {
		resMsg += fmt.Sprintf("%s: %s\n", cat, sum)
	}
	return resMsg, nil
}

var periodRegex = regexp.MustCompile(`\s*(week|month|year)\s*`)

func parsingReport(clientTxt string) (string, error) {
	if !periodRegex.MatchString(clientTxt) {
		Logger.Error("Couldn't parse client request")
		return "", errors.New("Couldn't parse client request")
	}
	groups := periodRegex.FindStringSubmatch(clientTxt)
	return groups[1], nil
}

func (m *Model) reportExpenses(ctx context.Context, msg Message) string {
	span, ctx := opentracing.StartSpanFromContext(ctx, "reportExpenses")
	defer span.Finish()

	period, err := parsingReport(msg.Text[len("/report"):])
	if err != nil {
		return parseErrMsg
	}

	key := fmt.Sprintf("%s%s%s",
		strconv.FormatInt(msg.UserID, 10),
		time.Now().Format("2006-01-02"),
		getSubDate(period).Format("2006-01-02"))
	val := m.cacheLRU.Get(key)
	if val != nil {
		log.Println("Get data from cache")
		return val.(string)
	}

	txtMsg, err := m.processingReport(ctx, msg.UserID, period)
	if err != nil {
		return internalErrMsg
	}
	m.cacheLRU.Add(key, txtMsg)
	return txtMsg
}

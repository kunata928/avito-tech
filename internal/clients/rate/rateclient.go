package rate

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

var historicalURL = "https://openexchangerates.org/api/historical/%s.json?app_id=%s"

//var latestURL = "https://openexchangerates.org/api/latest.json?app_id=%s"

type ClientRate struct {
	token  string
	client *http.Client
	//data *database.Rates
}

type TokenGetter interface {
	TokenRate() string
}

type Data struct {
	Base       string
	Disclaimer string
	License    string
	Timestamp  int64
	Rates      map[string]decimal.Decimal
}

func NewRateClient(tokenGetter TokenGetter) (*ClientRate, error) {
	return &ClientRate{
		token:  tokenGetter.TokenRate(),
		client: &http.Client{Timeout: 10 * time.Second},
	}, nil
}

func (c *ClientRate) GetRateDate(ctx context.Context, date time.Time) (*Data, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	url := fmt.Sprintf(historicalURL, date.Format("2006-01-02"), c.token)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.Error("Error request", zap.Error(err))
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		logger.Error("Error request", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get a list of currencies on the date")
	}
	defer resp.Body.Close()

	data := Data{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		logger.Error("Error decode response Rate API", zap.Error(err))
		return nil, err
	}

	logger.Debug("Get response from Rate API")
	return &data, nil
}

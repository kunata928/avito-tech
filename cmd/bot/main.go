package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shopspring/decimal"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/cache"
	db2 "gitlab.ozon.dev/kunata928/telegramBot/internal/clients/db"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/metrics"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gitlab.ozon.dev/kunata928/telegramBot/internal/clients/rate"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/clients/tg"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/config"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/database"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/logger"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/model/messages"
)

func main() {
	// ----- context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// ----- logger AND tracing
	logger.Info("Starting app...") //, zap.Int("port", 8080)

	// ----- metrics
	metrics.CommandStartTotal.Inc()
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			logger.Fatal("Err metrics")
		}
	}()

	// ----- tracing

	// ----- config
	cfg, err := config.New()
	if err != nil {
		logger.Fatal("Config init failed:", zap.Error(err))
		//log.Fatal("config init failed:", err)
	}

	// ----- DB client
	dbclient, _ := db2.NewDBClient(cfg)
	db, err := sql.Open("postgres", dbclient.Token)
	if err != nil {
		logger.Fatal("Open DB error:", zap.Error(err))
	}
	storage := database.NewStorage(db)
	// ----- Cache LRU
	cacheLRU := cache.NewLRU(1000)

	// ----- TG client
	tgClient, err := tg.New(cfg)
	if err != nil {
		logger.Fatal("Tg client init failed:", zap.Error(err))
		//log.Fatal("tg client init failed:", err)
	}

	// ----- Rate Client
	rateClient, err := rate.NewRateClient(cfg)
	if err != nil {
		logger.Fatal("Rate client init failed:", zap.Error(err))
		//log.Fatal("tg client init failed:", err)
	}

	// ----- MSG union clients
	msgModel := messages.New(tgClient, storage, rateClient, cacheLRU)

	// ----- Workers
	RateWorker(ctx, rateClient, storage)
	tgClient.ListenUpdates(ctx, msgModel)

	// ----- Close
	defer db.Close()
	defer logger.Info("Close DB Connection") //, zap.Int("port", 8080)

}

func RateWorker(ctx context.Context, rateClient *rate.ClientRate, storage *database.Storage) {
	ticker := time.NewTicker(time.Second * 60)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("stopped receiving exchange rates")
				return
			case <-ticker.C:
				select {
				case <-ctx.Done():
					log.Println("stopped receiving exchange rates")
					return
				default:
					data, err := rateClient.GetRateDate(ctx, time.Now())
					if err != nil {
						log.Println(err)
					} else {
						if !data.Rates["RUB"].Equal(decimal.NewFromInt(0)) {
							rates := database.Rates{}
							rates[database.RUB] = decimal.NewFromInt(1)
							rates[database.USD] = decimal.NewFromInt(1).Div(data.Rates["RUB"])
							rates[database.EUR] = rates[database.USD].Mul(data.Rates["EUR"])
							rates[database.CNY] = rates[database.USD].Mul(data.Rates["CNY"])
							if err := storage.SetRate(ctx, &rates, time.Now()); err != nil {
								log.Println(err, ":(")
							}
						}
					}
				}
			}
		}
	}()
}

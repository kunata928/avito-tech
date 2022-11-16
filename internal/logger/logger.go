package logger

import (
	"flag"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"log"
)

var (
	develMode   = flag.Bool("devel", false, "development mode")
	serviceName = flag.String("service", "bot", "the name of our service")
)

var (
	logger *zap.Logger
)

func initTracing(logger *zap.Logger) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(*serviceName)
	if err != nil {
		logger.Fatal("Cannot init tracing", zap.Error(err))
	}
}

func init() { //logger *zap.Logger
	var err error

	if *develMode {
		logger, err = zap.NewDevelopment()
	} else {
		cfg := zap.NewProductionConfig()
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		//cfg.Level.SetLevel()
		logger, err = cfg.Build()
	}

	if err != nil {
		log.Fatal("Cannot init zap.", err)
	}

	initTracing(logger)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

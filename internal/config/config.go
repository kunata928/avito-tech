package config

import (
	"gitlab.ozon.dev/kunata928/telegramBot/internal/logger"
	"go.uber.org/zap"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const configFile = "data/config.yaml"

type Config struct {
	TokenTg   string `yaml:"tgtoken"`
	TokenRate string `yaml:"ratetoken"`
	TokenDB   string `yaml:"dbtoken"`
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		logger.Error("Reading config file err", zap.Error(err))
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &s.config)
	if err != nil {
		logger.Error("Parsing yaml err", zap.Error(err))
		return nil, errors.Wrap(err, "parsing yaml")
	}

	return s, nil
}

func (s *Service) TokenTg() string {
	return s.config.TokenTg
}

func (s *Service) TokenRate() string {
	return s.config.TokenRate
}

func (s *Service) TokenDB() string {
	return s.config.TokenDB
}

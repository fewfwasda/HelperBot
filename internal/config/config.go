package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var configPath = ".env"

type (
	Config struct {
		DBConfig  DBConfig
		BotConfig BotConfig
	}
	DBConfig struct {
		DBUrl string `env:"DATABASE_URL"`
	}
	BotConfig struct {
		BotToken         string `env:"TOKEN_TELEGRAM_BOT"`
		UpdateBotTimeout int    `env:"UPDATE_BOT_TIMEOUT"`
	}
)

func New() (*Config, error) {
	conf := &Config{}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s is not exist", configPath)
	}
	err := cleanenv.ReadConfig(configPath, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

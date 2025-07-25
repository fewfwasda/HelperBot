package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var configPath = ".env"

type (
	Config struct {
		DBConfig  DBConfig
		BotConfig BotConfig
	}
	DBConfig struct {
		DBUrl        string        `env:"DATABASE_URL"`
		MaxConns     int32         `env:"MAX_CONNS"`
		MinConns     int32         `env:"MIN_CONNS"`
		MaxConnsIdle time.Duration `env:"MAX_CONNS_IDLE"`
	}
	BotConfig struct {
		BotToken         string `env:"TOKEN_TELEGRAM_BOT"`
		UpdateBotTimeout int    `env:"UPDATE_BOT_TIMEOUT"`
	}
)

func New() (*Config, error) {
	conf := &Config{}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s is not exist", configPath)
	}
	err := cleanenv.ReadConfig(configPath, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

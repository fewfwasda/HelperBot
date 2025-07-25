package bot

import (
	"HelperBot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Api           tgbotapi.BotAPI
	UpdateTimeout int
}

func NewBot(cfg config.BotConfig) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}
	return &Bot{
		Api:           *bot,
		UpdateTimeout: cfg.UpdateBotTimeout,
	}, nil
}
func (b Bot) GetUpdateConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = b.UpdateTimeout
	return updateConfig
}

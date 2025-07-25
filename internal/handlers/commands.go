package handlers

import (
	"HelperBot/pkg/bot"
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) Command(ctx context.Context, bot bot.Bot, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
	}

}

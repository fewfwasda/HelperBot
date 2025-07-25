package handlers

import (
	handlers "HelperBot/internal/handlers/start"
	"HelperBot/internal/services"
	"HelperBot/pkg/bot"
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler interface {
	HandleCommand(ctx context.Context, bot bot.Bot, update tgbotapi.Update) error
}

type Handlers struct {
	StartHandler *handlers.StartHandler
}

func NewHandler(userService services.UserService) *Handlers {
	return &Handlers{
		StartHandler: handlers.NewStartHandler(userService),
	}
}

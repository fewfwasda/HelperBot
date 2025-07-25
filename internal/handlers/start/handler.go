package handlers

import (
	"HelperBot/internal/services"
	"HelperBot/pkg/bot"
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
	userService services.UserService
}

func NewStartHandler(userService services.UserService) *StartHandler {
	return &StartHandler{
		userService: userService,
	}
}

func (h *StartHandler) HandleCommand(ctx context.Context, bot bot.Bot, update tgbotapi.Update) error {
	panic("asd")
}

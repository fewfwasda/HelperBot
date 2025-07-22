package startbot

import (
	message "HelperBot"
	"HelperBot/init/bot"
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(b *bot.Bot, db *sql.DB) error {
	updateCfg := tgbotapi.NewUpdate(0)
	updateCfg.Timeout = b.UpdateTimeout
	updates := b.Api.GetUpdatesChan(updateCfg)

	for update := range updates {
		if update.Message != nil {
			text, err := message.GetMessageContext(db, "WelcomeText")
			if err != nil {
				text = "такой текст не найден"
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			b.Api.Send(msg)
		}
	}
	return nil
}

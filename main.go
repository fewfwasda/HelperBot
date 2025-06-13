package main

import (
	taskmanager "HelperBot/taskManager"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load env file %v", err)
	}
	botToken := os.Getenv("TOKEN_TELEGRAM_BOT")
	if botToken == "" {
		log.Fatalf("Token is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true
	log.Printf("Autorization on account %v", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	tasks := make(map[int][]string)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		userID := update.Message.From.ID
		text := update.Message.Text

		switch {
		case strings.HasPrefix(text, "/list"):
			userTasks := tasks[int(userID)]
			response := taskmanager.ShowTaskList(userTasks)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			bot.Send(msg)
		case strings.HasPrefix(text, "/clear"):
			taskmanager.ClearTasks(tasks, int(userID))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Все ваши удалены")
			bot.Send(msg)
		case strings.HasPrefix(strings.ToLower(text), "сделано"):
			parts := strings.Fields(text)
			if len(parts) != 2 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка. Используйте формат сделано <номер задачи>")
				bot.Send(msg)
				continue
			}
			number, err := strconv.Atoi(parts[1])
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверный номер задачи")
				bot.Send(msg)
				continue
			}
			confirmation, err := taskmanager.RemoveTasks(tasks, int(userID), number)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				bot.Send(msg)
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, confirmation)
			bot.Send(msg)
		default:
			taskmanager.AddTask(tasks, int(userID), text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Задача добавлена: %v", text))
			bot.Send(msg)
		}
	}

}

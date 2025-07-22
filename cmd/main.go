package main

import (
	"HelperBot/init/bot"
	"HelperBot/init/db"
	"HelperBot/internal/config"

	startbot "HelperBot/start_bot"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err.Error())
	}
	bot, err := bot.NewBot(cfg.BotConfig)
	if err != nil {
		panic(err.Error())
	}
	db, pool, err := db.NewDatabase(cfg.DBConfig)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	defer pool.Close()
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	startbot.Start(bot, db)
}

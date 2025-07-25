package app

import (
	"HelperBot/internal/config"
	"HelperBot/pkg/bot"
	"HelperBot/pkg/db"
	"context"
	"time"

	"github.com/pressly/goose"
)

type App struct {
	bot       *bot.Bot
	db        db.Database
	appConfig *config.Config
	ctx       context.Context
}

func NewApp(cfg *config.Config) *App {
	ctx := context.Background()

	app := &App{
		appConfig: cfg,
		ctx:       ctx,
	}

	app.mustCreateBot()
	app.mustCteateDB()
	app.mustCreateTabe()

	return app
}

func (a *App) Run() {
	go a.startBot()
	<-a.ctx.Done()
	time.Sleep(2 * time.Second)
}

func (a *App) mustCteateDB() {
	db, err := db.NewDatabase(a.ctx, a.appConfig.DBConfig)
	if err != nil {
		panic(err.Error())
	}

	a.db.DB = db.DB
	a.db.Pool = db.Pool
}

func (a *App) mustCreateBot() {
	bot, err := bot.NewBot(a.appConfig.BotConfig)
	if err != nil {
		panic(err.Error())
	}

	a.bot = bot
}

func (a *App) mustCreateTabe() {
	if err := goose.Up(a.db.DB, "migrations"); err != nil {
		panic(err)
	}
}

func (a *App) startBot() {
	updateCfg := a.bot.GetUpdateConfig()

	updates := a.bot.Api.GetUpdatesChan(updateCfg)

	for update := range updates {
		if update.Message.IsCommand() {

		}
	}
}

func (a *App) Close() {
	if a.db.DB != nil {
		a.db.DB.Close()
	}
	if a.db.Pool != nil {
		a.db.Pool.Close()
	}
}

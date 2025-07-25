package main

import (
	"HelperBot/cmd/app"
	"HelperBot/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		panic(err.Error())
	}
	botApp := app.NewApp(cfg)

	defer botApp.Close()

	botApp.Run()
}

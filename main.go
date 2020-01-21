package main

import (
	"flag"

	"speedy/bot"
	"speedy/config"
	"speedy/db"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "config.toml", "Configuration file path")
	flag.Parse()

	config.Init(configFilePath)
	config.InitLogging()
	db.Init()
	defer db.Close()

	bot.Start(config.Env.Telegram.Token)
}

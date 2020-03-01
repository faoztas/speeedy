package main

import (
	"flag"

	"github.com/speeedy/bot"
	"github.com/speeedy/config"
	"github.com/speeedy/db"
	"github.com/speeedy/job"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "config.toml", "Configuration file path")
	flag.Parse()

	config.Init(configFilePath)
	config.InitLogging()

	db.Init()
	db.Migrate(db.DB)
	defer db.Close()

	go job.Jobs()

	bot.Start(config.Env.Telegram.Token)
}

package bot

import (
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func Start(token string) {
	b, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	h := Handler{Bot: b}
	b.Handle("/start", h.Start)
	b.Handle("/list", h.ServerLists)
	b.Handle("/test", h.Test)
	b.Handle("/help", h.Help)
	// b.Handle(telebot.OnText, h.Text)


	logrus.Info("Telegram bot starting")
	b.Start()
}
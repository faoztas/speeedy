package utils

import (
	"errors"
	"runtime/debug"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"

	"github.com/speeedy/db"
	"github.com/speeedy/models"
)

func User(m telebot.Message) {
	user := models.User{
		SenderID:     m.Sender.ID,
		UserName:     m.Sender.Username,
		FirstName:    m.Sender.FirstName,
		LastName:     m.Sender.LastName,
		LanguageCode: m.Sender.LanguageCode,
		IsBot:        m.Sender.IsBot,
	}
	if err := db.DB.Where("sender_id = ?", m.Sender.ID).Assign(user).FirstOrCreate(&user).Error; err != nil {
		logrus.Error(err)
	}
}

func RunRoutineForever(fn func(), delay time.Duration) {
	go func() {
		defer retryRecover(fn, delay)
		fn()
	}()
}

func retryRecover(fn func(), delay time.Duration) {
	if r := recover(); r != nil {
		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("unknown panic")
		}

		logrus.WithFields(logrus.Fields{
			"trace": string(debug.Stack()),
		}).WithError(err).Error("RunRoutineForever recovered")
	}

	logrus.Infof("Task will be run after %v seconds", delay.Seconds())
	time.Sleep(delay)
	RunRoutineForever(fn, delay)
}


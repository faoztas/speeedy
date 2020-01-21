package config

import (
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogging() {
	if Env.Logging.Level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if Env.Logging.Level == "info" {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if Env.Logging.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else if Env.Logging.Format == "text" {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	// For catching panic or fatal errors, we are waiting before exist.
	logrus.RegisterExitHandler(func() {
		time.Sleep(3 * time.Second)
	})
}

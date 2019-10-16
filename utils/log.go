package utils

import (
	logrus "github.com/sirupsen/logrus"
)

func init() {
	initLog()
}

var Log *logrus.Logger

func initLog() {
	Log = logrus.New()
	Log.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "01-02 15:04:05.000000",
	}
	Log.WithFields(logrus.Fields{"package": "moleculer-go"})
}

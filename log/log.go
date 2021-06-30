package log

import (
	"github.com/sirupsen/logrus"
	"github.com/xinxuwang/ExchangeX/log/hooks"
	"io"
	"os"
)

func Init(level logrus.Level, logPath, serverName string) (err error) {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logfile, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	mw := io.MultiWriter(os.Stdout, logfile)
	logrus.SetOutput(mw)
	logrus.SetLevel(level)
	logrus.AddHook(hooks.NewCallStackHook())
	logrus.AddHook(hooks.NewServerHook(serverName))

	logrus.WithFields(logrus.Fields{"serverName": serverName})
	return
}

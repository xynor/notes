package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.AddHook(DefaultFieldHook{})
	logfile, _ := os.OpenFile("./app.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	logrus.SetOutput(logfile) //默认为os.stderr
}

func GetC() *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"ABC": "dfsfds",
		"CD":  "vdfd",
	})
}

type DefaultFieldHook struct {
}

func (hook DefaultFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = "MyAppName"
	return nil
}

func (hook DefaultFieldHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

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
	// rotate writer
	//rotateLog, err := rotatelogs.New(
	//	logPath+".%Y%m%d",
	//	rotatelogs.WithLinkName(logPath),
	//	rotatelogs.WithRotationTime(24*time.Hour),
	//	rotatelogs.WithMaxAge(10*365*24*time.Hour), //10 years
	//)
	//if err != nil {
	//	return
	//}

	//mw := io.MultiWriter(os.Stdout, rotateLog)
	//Use stdout for current dev
	mw := io.MultiWriter(os.Stdout)
	logrus.SetOutput(mw)
	logrus.SetLevel(level)
	logrus.AddHook(hooks.NewCallStackHook())
	logrus.AddHook(hooks.NewServerHook(serverName))

	return
}

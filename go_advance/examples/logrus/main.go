package main

import (
	"./log"
	"github.com/sirupsen/logrus"
)

//方式一：logrus函数（最终调用的是logrus.StandardLogger默认实例方法）
func main() {
	log.GetC().Infof("ABD")
	logrus.Info("CDX")
	logrus.Error("EEEE")
	log.GetC().WithField("sdfd", "dfdsf").Infof("with")
}

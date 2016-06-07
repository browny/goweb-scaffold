package logger

import (
	"io/ioutil"

	"goweb-scaffold/config"

	"github.com/cihub/seelog"
)

var logger seelog.LoggerInterface

// SetupLogger setups customized logger
func SetupLogger() {
	seelogConf, _ := ioutil.ReadAll(config.LoadAsset("/config/seelog.xml"))
	l, _ := seelog.LoggerFromConfigAsBytes(seelogConf)
	l.SetAdditionalStackDepth(1)

	logger = l
}

func Tracef(format string, params ...interface{}) {
	logger.Tracef(format, params...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Debugf(format string, params ...interface{}) {
	logger.Debugf(format, params...)
}

func Infof(format string, params ...interface{}) {
	logger.Infof(format, params...)
}

func Warnf(format string, params ...interface{}) {
	logger.Warnf(format, params...)
}

func Errorf(format string, params ...interface{}) {
	logger.Errorf(format, params...)
}

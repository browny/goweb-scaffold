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

// Tracef wraps seelog.Tracef
func Tracef(format string, params ...interface{}) {
	logger.Tracef(format, params...)
}

// Debug wraps seelog.Debug
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Debugf wraps seelog.Debugf
func Debugf(format string, params ...interface{}) {
	logger.Debugf(format, params...)
}

// Infof wraps seelog.Infof
func Infof(format string, params ...interface{}) {
	logger.Infof(format, params...)
}

// Warnf wraps seelog.Warnf
func Warnf(format string, params ...interface{}) {
	logger.Warnf(format, params...)
}

// Errorf wraps seelog.Errorf
func Errorf(format string, params ...interface{}) {
	logger.Errorf(format, params...)
}

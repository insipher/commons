package utils

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
	logMiddleware "github.com/neko-neko/echo-logrus/v2"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
)

// Log serves as a wrapper around the underlying logging implementation to provide a consistent logging framework independent of underlying logger
var Log *log.MyLogger

// LogMiddleware serves as a wrapper around the underlying middleware logging implementation to provide a consistent logging framework independent of underlying logger
var LogMiddleware echo.MiddlewareFunc

// LoggerInit initializes the logger
func LoggerInit(lvl string) {

	// configure application logger
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(getEchoLogLevel(lvl))
	log.Logger().SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	log.Info("Logger initialized.")

	Log = log.Logger()
	LogMiddleware = logMiddleware.Logger()
}

// getEchoLogLevel gets the log level from the configuration
func getEchoLogLevel(lvl string) echoLog.Lvl {
	var level echoLog.Lvl

	switch lvl {
	case "debug":
		level = echoLog.DEBUG
	case "info":
		level = echoLog.INFO
	case "warn":
		level = echoLog.WARN
	case "error":
		level = echoLog.ERROR
	}

	return level
}

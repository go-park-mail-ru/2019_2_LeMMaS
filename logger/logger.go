package logger

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo"
	"os"
)

var e echo.Echo

func Init(echoInstance echo.Echo) {
	e = echoInstance
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		e.Logger.Error("error connecting to Sentry", err)
	}
}

func Error(err error) {
	e.Logger.Error(err)
	sentry.CaptureException(err)
}

func Errorf(format string, args ...interface{}) {
	e.Logger.Errorf(format, args...)
	sentry.CaptureException(fmt.Errorf(format, args...))
}

func Warn(err error) {
	e.Logger.Warn(err)
	sentry.CaptureMessage(err.Error())
}

func Warnf(format string, args ...interface{}) {
	e.Logger.Errorf(format, args...)
	sentry.CaptureMessage(fmt.Sprintf(format, args...))
}

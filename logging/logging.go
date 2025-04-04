package logging

import (
	"errors"
	"os"
	"runtime/debug"

	"github.com/cto-kit/service"

	"github.com/sirupsen/logrus"
)

var (
	logLevels = map[string]logrus.Level{
		"DEBUG":    logrus.DebugLevel,
		"INFO":     logrus.InfoLevel,
		"WARNING":  logrus.WarnLevel,
		"ERROR":    logrus.ErrorLevel,
		"CRITICAL": logrus.FatalLevel,
		"PANIC":    logrus.PanicLevel,
	}
	logLevel = logrus.InfoLevel
	appName  = ""
)

type hook struct{}

func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *hook) Fire(entry *logrus.Entry) error {

	for key, value := range entry.Data {
		if v, ok := value.(string); ok && v == "" {
			delete(entry.Data, key)
			continue
		}
	}

	return nil

	/* github copilot code
	if appName == "" {
		return nil
	}
	entry.Data["app"] = appName
	return nil
	*/
}

func Init(app string, opts ...Option) {
	opt := option{}
	for _, o := range opts {
		o(&opt)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap:    logrus.FieldMap{},
		PrettyPrint: opt.prettyPrint,
	})

	appName = app
	logrus.AddHook(&hook{})
}

func Logging(level logrus.Level, code string, message string, opts ...Option) {
	if logLevel < level {
		return
	}

	opt := option{
		service: service.Application,
	}

	for _, o := range opts {
		o(&opt)
	}

	file, line, _ := getFileAndLine()

	method := getMethod()
	host, _ := os.Hostname()
	fields := logrus.Fields{
		"service":  opt.service,
		"file":     file,
		"line":     line,
		"method":   method,
		"host":     host,
		"app":      appName,
		"code":     code,
		"identity": opt.identity,
		"ref":      opt.reference,
		"trace":    opt.errors,
		"id":       opt.referenceID,
		"versopm":  opt.appVersion,
	}

	if level == logrus.FatalLevel || opt.stackTrace {
		fields["stack"] = string(debug.Stack())
	}

	logrus.WithFields(fields).Log(level, message)
}

func Error(code string, message string, opts ...Option) {
	Logging(logrus.ErrorLevel, code, message, opts...)
}

func Warning(code string, message string, opts ...Option) {
	Logging(logrus.WarnLevel, code, message, opts...)
}

func Info(code string, message string, opts ...Option) {
	Logging(logrus.InfoLevel, code, message, opts...)
}

func Debug(code string, message string, opts ...Option) {
	Logging(logrus.DebugLevel, code, message, opts...)
}

func Critical(code string, message string, opts ...Option) {
	Logging(logrus.FatalLevel, code, message, opts...)
}

func SetLevel(level string) error {
	if numLevel, ok := logLevels[level]; ok {
		logrus.SetLevel(numLevel)
		logLevel = numLevel
		return nil
	}
	return errors.New("invalid log level" + level)
}

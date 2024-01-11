package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/blawhi2435/shanjuku-backend/environment"
	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
)

type LoggerService struct {
	Log      *logrus.Logger
	Path     string
	FileName string
	IsTest   bool
}

func ProvideLogger() (*LoggerService, error) {
	logger := logrus.New()
	level, err := logrus.ParseLevel(environment.Setting.Logger.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)
	logger.Formatter = &logrus.JSONFormatter{}
	if _, err := os.Stat(environment.Setting.Logger.Path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(environment.Setting.Logger.Path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return &LoggerService{
		Log:      logger,
		Path:     environment.Setting.Logger.Path,
		FileName: environment.Setting.Logger.FileName,
	}, nil
}

type Labels struct {
	RequestID string `json:"request_id" structs:"request_id"`
	Source    string `json:"source" structs:"source"`
}

type DefaultLogEntry struct {
	Labels   Labels `json:"labels" structs:"logging.googleapis.com/labels"`
	Severity string `json:"severity" structs:"logging.googleapis.com/severity"`
}

type HTTPRequestLogEntry struct {
	HttpRequest HttpRequest `json:"http_request" structs:"logging.googleapis.com/httpRequest"`
	Labels      Labels      `json:"labels" structs:"logging.googleapis.com/labels"`
	Severity    string      `json:"severity" structs:"logging.googleapis.com/severity"`
	RequestBody any         `json:"request_body" structs:"requestBody"`
}

type HTTPResponseLogEntry struct {
	HttpRequest  HttpRequest `json:"http_request" structs:"logging.googleapis.com/httpRequest"`
	Labels       Labels      `json:"labels" structs:"logging.googleapis.com/labels"`
	Severity     string      `json:"severity" structs:"logging.googleapis.com/severity"`
	UserID       int64       `json:"user_id" structs:"userId"`
	ResponseBody any         `json:"response_body" structs:"responseBody"`
}

type HttpRequest struct {
	RequestMethod string `json:"request_method" structs:"requestMethod"`
	RequestURL    string `json:"request_url" structs:"requestUrl"`
	Status        int    `json:"status" structs:"status"`
	UserAgent     string `json:"user_agent" structs:"userAgent"`
	RemoteIP      string `json:"remote_ip" structs:"remoteIp"`
	Latency       string `json:"latency" structs:"latency"`
	Protocol      string `json:"protocol" structs:"protocol"`
}

func (l *LoggerService) CombineErrorMessage(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func (l *LoggerService) GetLoggerFields(logEntry any) logrus.Fields {

	return structs.Map(logEntry)
}

func (l *LoggerService) LogFile(ctx context.Context, level logrus.Level, logEntry any, message string) {

	if l.IsTest {
		return
	}

	fields := l.GetLoggerFields(logEntry)

	filePath := l.Path + l.FileName + "-" + time.Now().Format("20060102") + ".log"
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	l.Log.SetOutput(f)

	switch level {
	case logrus.FatalLevel:
		l.Log.WithFields(fields).Fatal(message)
	case logrus.PanicLevel:
		l.Log.WithFields(fields).Panicf(message)
	case logrus.ErrorLevel:
		l.Log.WithFields(fields).Error(message)
	case logrus.WarnLevel:
		l.Log.WithFields(fields).Warn(message)
	case logrus.InfoLevel:
		l.Log.WithFields(fields).Info(message)
	case logrus.DebugLevel:
		l.Log.WithFields(fields).Debug(message)
	default:
		l.Log.WithFields(fields).Info(message)
	}
	return
}

func (l *LoggerService) GetLogEntrySeverity(level int) string {
	switch level {
	case int(logrus.DebugLevel):
		return "DEBUG"
	case int(logrus.InfoLevel):
		return "INFO"
	case int(logrus.WarnLevel):
		return "WARN"
	case int(logrus.ErrorLevel):
		return "ERROR"
	case int(logrus.PanicLevel):
		return "CRITICAL"
	case int(logrus.FatalLevel):
		return "ALERT"
	default:
		return "DEFAULT"
	}
}

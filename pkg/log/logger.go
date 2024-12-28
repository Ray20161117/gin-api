/**
 * log日志
 */
package log

import (
	config "gin-api/config/yaml_config"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var logToFile *logrus.Logger

// 初始化日志文件名
func init() {
	loggerFile := filepath.Join(config.Cfg.Log.Path, config.Cfg.Log.Name)
	if config.Cfg.Log.Model == "file" {
		logToFile = initializeFileLogger(loggerFile)
	}
}

// Log 方法调用
func Log() *logrus.Logger {
	if config.Cfg.Log.Model == "file" {
		if logToFile == nil {
			loggerFile := filepath.Join(config.Cfg.Log.Path, config.Cfg.Log.Name)
			logToFile = initializeFileLogger(loggerFile)
		}
		return logToFile
	}

	if log == nil {
		log = logrus.New()
		log.Out = os.Stdout
		log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}
		log.SetLevel(logrus.DebugLevel)
	}
	return log
}

// 初始化文件日志记录器
func initializeFileLogger(file string) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		file+"_%Y%m%d.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logger.Errorf("Failed to create rotatelogs: %v", err)
		return nil
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(lfHook)
	return logger
}

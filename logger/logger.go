package logger

import (
	"path"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type Logger struct {
	path          string
	filename      string
	maxSize       int
	maxBackups    int
	maxAge        int
	logLevel      logger.LogLevel
	slowThreshold time.Duration
}

func New(opts ...Option) *Logger {
	l := &Logger{
		path:          "./",
		filename:      "gorm.log",
		maxSize:       100,
		maxBackups:    5,
		maxAge:        3,
		logLevel:      logger.Error,
		slowThreshold: time.Second,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (l *Logger) Logger() logger.Interface {
	writer := &lumberjack.Logger{
		Filename:   path.Join(l.path, l.filename),
		MaxSize:    l.maxSize,    // 文件大小限制,单位:MB
		MaxBackups: l.maxBackups, // 最大保留日志文件数量
		MaxAge:     l.maxAge,     // 日志文件保留天数
		Compress:   false,        // 是否压缩处理
		LocalTime:  true,
	}

	loggerSplit := logrus.New()
	loggerSplit.SetOutput(writer)
	loggerSplit.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.DateTime})

	var gormLogger logger.Interface

	gormLogger = logger.New(loggerSplit, logger.Config{
		SlowThreshold:             l.slowThreshold, // Slow SQL threshold
		LogLevel:                  l.logLevel,      // Log level
		IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      false,           // Don't include params in the SQL log
		Colorful:                  false,           // Disable color
	})

	return gormLogger
}

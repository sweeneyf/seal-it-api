package logfactory

import (
	"fmt"

	"github.com/sweeneyf/seal-it-api/entity"
	"github.com/sweeneyf/seal-it-api/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	zap *zap.Logger
}

func CreateLogger(logLevel logger.LogLevel, logPath string) *zapLogger {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout", logPath},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapLogLevel(logLevel)),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	zapLog, err := logConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return &zapLogger{
		zap: zapLog,
	}
}

func zapField(kvPair entity.KvPair) zap.Field {
	return zap.Any(kvPair.Key, kvPair.Value)
}

// convert application log level to zaploglevel, if not found default to err
func zapLogLevel(ll logger.LogLevel) zapcore.Level {
	//default to error level
	var level zapcore.Level = zap.ErrorLevel
	switch ll {
	case logger.LogLevel(logger.Debug):
		level = zap.DebugLevel
	case logger.LogLevel(logger.Info):
		level = zap.InfoLevel
	case logger.LogLevel(logger.Warn):
		level = zap.WarnLevel
	case logger.LogLevel(logger.Error):
		level = zap.ErrorLevel
	}
	return level
}

func (l *zapLogger) Debug(msg string, tags ...(entity.KvPair)) {
	zapTags := make([]zap.Field, len(tags))
	for i := 0; i < len(tags); i++ {
		zapTags[i] = zapField(tags[i])
	}
	l.zap.Debug(msg, zapTags...)
	l.zap.Sync()
}

func (l *zapLogger) Info(msg string, tags ...(entity.KvPair)) {
	zapTags := make([]zap.Field, len(tags))
	for i := 0; i < len(tags); i++ {
		zapTags[i] = zapField(tags[i])
	}
	l.zap.Info(msg, zapTags...)
	l.zap.Sync()
}

func (l *zapLogger) Error(msg string, err error, tags ...(entity.KvPair)) {
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	zapTags := make([]zap.Field, len(tags))
	for i := 0; i < len(tags); i++ {
		zapTags[i] = zapField(tags[i])
	}
	l.zap.Error(msg, zapTags...)
	l.zap.Sync()
}

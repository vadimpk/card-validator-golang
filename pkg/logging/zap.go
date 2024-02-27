package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

var _ Logger = (*zapLogger)(nil)

func New(level string) *zapLogger {
	var l zapcore.Level
	l, err := zapcore.ParseLevel(level)
	if err != nil {
		l = zap.InfoLevel
	}

	config := zap.Config{
		Development:      false,
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(l),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			EncodeDuration: zapcore.SecondsDurationEncoder,
			LevelKey:       "severity",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			CallerKey:      "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,
			TimeKey:        "timestamp",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			NameKey:        "name",
			EncodeName:     zapcore.FullNameEncoder,
			MessageKey:     "message",
			StacktraceKey:  "",
			LineEnding:     "\n",
		},
	}

	logger, _ := config.Build()

	return &zapLogger{
		logger: logger.Sugar(),
	}
}

func (l *zapLogger) Named(name string) Logger {
	return &zapLogger{
		logger: l.logger.Named(name),
	}
}

func (l *zapLogger) With(args ...interface{}) Logger {
	return &zapLogger{
		logger: l.logger.With(args...),
	}
}
func (l *zapLogger) Debug(message string, args ...interface{}) {
	l.logger.Debugw(message, args...)
}

func (l *zapLogger) Info(message string, args ...interface{}) {
	l.logger.Infow(message, args...)
}

func (l *zapLogger) Warn(message string, args ...interface{}) {
	l.logger.Warnw(message, args...)
}

func (l *zapLogger) Error(message string, args ...interface{}) {
	l.logger.Errorw(message, args...)
}

func (l *zapLogger) Fatal(message string, args ...interface{}) {
	l.logger.Fatalw(message, args...)
	os.Exit(1)
}

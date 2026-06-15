package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/sagargaikwad2000/logger/constant"
)

type ILogger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type Logger struct {
	logger *slog.Logger
}

func New(file *os.File, level string, enableFileSource bool) ILogger {
	var output io.Writer

	if file != nil {
		output = io.MultiWriter(file, os.Stdout)
	} else {
		output = os.Stdout
	}

	var slogLevel slog.Level

	switch strings.ToLower(strings.TrimSpace(level)) {
	case constant.DEBUG:
		slogLevel = slog.LevelDebug

	case constant.WARN:
		slogLevel = slog.LevelWarn

	case constant.ERROR:
		slogLevel = slog.LevelError

	default:
		slogLevel = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(output, &slog.HandlerOptions{
		Level:     slogLevel,
		AddSource: enableFileSource,
	})

	return &Logger{
		logger: slog.New(handler),
	}
}

// Debug

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type ILogger interface {
	Debug(string)
	Debugf(string, ...interface{})
	Info(string)
	Infof(string, ...interface{})
	Warn(string)
	Warnf(string, ...interface{})
	Error(string)
	Errorf(string, ...interface{})
}

type Logger struct {
	debug *slog.Logger
	info  *slog.Logger
	warn  *slog.Logger
	err   *slog.Logger
}

/*
Create new instance of Logger
Parameter : file or nil
*/
func New(file *os.File) ILogger {
	var output io.Writer
	if file != nil {
		output = io.MultiWriter(file, os.Stdout)
	} else {
		output = os.Stdout
	}

	debugLogger := slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{Level: slog.LevelDebug}))
	infoLogger := slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{Level: slog.LevelInfo}))
	warnLogger := slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{Level: slog.LevelWarn}))
	errorLogger := slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{Level: slog.LevelError}))

	return &Logger{debugLogger, infoLogger, warnLogger, errorLogger}
}

// Debug
func (l *Logger) Debug(msg string) {
	l.debug.Debug(msg, slog.String("file", getCallerFile()))
}

func (l *Logger) Debugf(s string, args ...interface{}) {
	output := fmt.Sprintf(s, args...)
	l.debug.Debug(output, slog.String("file", getCallerFile()))
}

// Info
func (l *Logger) Info(msg string) {
	l.info.Info(msg)
}

func (l *Logger) Infof(s string, args ...interface{}) {
	output := fmt.Sprintf(s, args...)
	l.info.Info(output)
}

// Warn
func (l *Logger) Warn(msg string) {
	l.warn.Warn(msg)
}

func (l *Logger) Warnf(s string, args ...interface{}) {
	output := fmt.Sprintf(s, args...)
	l.warn.Warn(output)
}

// Error
func (l *Logger) Error(msg string) {
	l.err.Error(msg)
}

func (l *Logger) Errorf(s string, args ...interface{}) {
	output := fmt.Sprintf(s, args...)
	l.err.Error(output)
}

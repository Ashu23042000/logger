package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/Ashu23042000/logger/constant"
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
	level string
}

/*
New creates a new instance of ILogger with specified log level and output file.

Parameters:
  - file: An optional *os.File to write log output to. If nil, logs are only written to standard output.
  - level: A string representing the desired log level. This should be one of "debug", "info", "warn", or "error".

Returns:
  - ILogger: An instance of a Logger that supports different log levels (debug, info, warn, error). The log messages
             are formatted in JSON and can be written to the specified file and standard output, or just standard output
             if no file is provided.
*/
func New(file *os.File, level string) ILogger {
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
	level = strings.TrimSpace(strings.ToLower(level))

	return &Logger{debugLogger, infoLogger, warnLogger, errorLogger, level}
}

// Info
func (l *Logger) Info(msg string) {
	l.info.Info(msg)
}

func (l *Logger) Infof(s string, args ...interface{}) {
	output := fmt.Sprintf(s, args...)
	l.info.Info(output)
}

// Debug
func (l *Logger) Debug(msg string) {
	if l.level == constant.DEBUG {
		l.debug.Debug(msg, slog.String("file", getCallerFile()))
	}
}

func (l *Logger) Debugf(s string, args ...interface{}) {
	if l.level == constant.DEBUG {
		output := fmt.Sprintf(s, args...)
		l.debug.Debug(output, slog.String("file", getCallerFile()))
	}
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

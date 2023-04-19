package pocketlog

import "fmt"

type Logger struct {
    threshold Level
}

func New(threshold Level) *Logger {
    return &Logger{
        threshold: threshold,
    }
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
}

// Info: formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
}

func (l *Logger) Errorf(format string, args ...any) {
}

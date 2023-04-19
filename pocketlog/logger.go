package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
    threshold Level
    output io.Writer
}

// New returns you a logger, ready to log at the required threshold
// The default output is Stdout.
func New(threshold Level, output io.Writer) *Logger {
    return &Logger{
        threshold: threshold,
        output: output,
    }
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
    // making sure we can safely write to the output
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelDebug {
        return 
    }

    // _, _ is a way to indicate to anyone reading the code that we're 
    // aware of the return value from the function call on the right 
    // but we don't need them at the moment.
    _, _ = fmt.Printf(format + "\n", args...)
}

// Info formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelInfo {
        return 
    }

    _, _ = fmt.Printf(format + "\n", args...)
}

// Error formats and prints a message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelError {
        return 
    }

    _, _ = fmt.Printf(format + "\n", args...)
}

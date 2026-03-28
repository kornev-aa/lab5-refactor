package logger

import (
    "fmt"
    "time"
)

type Logger struct{}

func New() *Logger {
    return &Logger{}
}

func (l *Logger) Info(msg string) {
    fmt.Println(l.format("INFO", msg, nil))
}

func (l *Logger) Debug(msg string) {
    fmt.Println(l.format("DEBUG", msg, nil))
}

func (l *Logger) Error(msg string, err error) {
    fmt.Println(l.format("ERROR", msg, err))
}

func (l *Logger) format(level, msg string, err error) string {
    timeStr := time.Now().Format(time.RFC3339)
    if err != nil {
        return fmt.Sprintf("[%s] %s - %s: %v", level, timeStr, msg, err)
    }
    return fmt.Sprintf("[%s] %s - %s", level, timeStr, msg)
}

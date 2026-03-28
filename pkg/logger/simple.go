package logger

import (
    "fmt"
    "time"
)

type SimpleLogger struct{}

func NewSimpleLogger() *SimpleLogger {
    return &SimpleLogger{}
}

func (l *SimpleLogger) Info(msg string) {
    fmt.Printf("[INFO] %s - %s\n", time.Now().Format("15:04:05"), msg)
}

func (l *SimpleLogger) Debug(msg string) {
    fmt.Printf("[DEBUG] %s - %s\n", time.Now().Format("15:04:05"), msg)
}

func (l *SimpleLogger) Error(msg string) {
    fmt.Printf("[ERROR] %s - %s\n", time.Now().Format("15:04:05"), msg)
}

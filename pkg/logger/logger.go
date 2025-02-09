package logger

import (
	"log"
	"log/slog"
	"os"
	"sync"
)

var (
	instance *Logger
	once     sync.Once
)

type Logger struct {
	loggers map[string]*slog.Logger
	mu      sync.Mutex
}

// GetLogger возвращает единственный экземпляр логгера
func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			loggers: make(map[string]*slog.Logger),
		}
	})
	return instance
}

// getFileLogger создает новый логгер для указанного файла
func (l *Logger) getFileLogger(logType string) *slog.Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	if logger, exists := l.loggers[logType]; exists {
		return logger
	}

	file, err := os.OpenFile("logs/"+logType+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))
	l.loggers[logType] = logger
	return logger
}

// Debug пишет debug-сообщение
func (l *Logger) Debug(msg string, args ...any) {
	l.getFileLogger("debug").Debug(msg, args...)
}

// Info пишет info-сообщение
func (l *Logger) Info(msg string, args ...any) {
	l.getFileLogger("info").Info(msg, args...)
}

// Error пишет error-сообщение
func (l *Logger) Error(msg string, args ...any) {
	l.getFileLogger("error").Error(msg, args...)
}

// Printf пишет сообщение в stdout
func (l *Logger) Printf(msg string, args ...any) {
	log.Printf(msg, args...)
}

// Fatalf пишет сообщение в stderr и завершает работу программы
func (l *Logger) Fatalf(msg string, args ...any) {
	log.Fatalf(msg, args...)
}

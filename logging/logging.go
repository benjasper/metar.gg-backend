package logging

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Debug(input string) {
	message := fmt.Sprintf("[DEBUG] %s\n", input)

	l.logger.Printf(message)
	log.Print(message)
}

func (l *Logger) Info(input string) {
	message := fmt.Sprintf("[INFO] %s\n", input)

	l.logger.Printf(message)
	log.Print(message)
}

func (l *Logger) Warn(input string) {
	message := fmt.Sprintf("[WARN] %s\n", input)

	l.logger.Printf(message)
	log.Print(message)
}

func (l *Logger) Error(input string) {
	message := fmt.Sprintf("[ERROR] %s\n", input)

	l.logger.Printf(message)
	log.Print(message)
}

func (l *Logger) Fatal(input error) {
	message := fmt.Sprintf("[FATAL] %s\n", input)

	l.logger.Printf(message)
	log.Fatal(message)
}

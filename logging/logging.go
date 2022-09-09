package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/axiomhq/axiom-go/axiom"
	"log"
	"os"
	"time"
)

const logFile = "./tmp/logs.log"

type Logger struct {
	logger *log.Logger
}

type Message struct {
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (m *Message) String() string {
	return fmt.Sprintf("%s [%s] %s\n", m.Time.Format("2006-01-02 15:04:05"), m.Type, m.Message)
}

func NewLogger() *Logger {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", 0)

	loggerObject := &Logger{
		logger: logger,
	}

	// Trigger upload every minute
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			loggerObject.uploadLog()
		}
	}()

	return loggerObject
}

func (l *Logger) Debug(input string) {
	m := Message{
		Type:    "DEBUG",
		Message: input,
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[DEBUG] %s\n", input)
	l.messageToJson(&m)

	log.Print(message)
}

func (l *Logger) Info(input string) {
	m := Message{
		Type:    "INFO",
		Message: input,
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[INFO] %s\n", input)
	l.messageToJson(&m)

	log.Print(message)
}

func (l *Logger) Warn(input string) {
	m := Message{
		Type:    "WARN",
		Message: input,
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[WARN] %s\n", input)
	l.messageToJson(&m)

	log.Print(message)
}

func (l *Logger) Error(input string) {
	m := Message{
		Type:    "ERROR",
		Message: input,
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[ERROR] %s\n", input)
	l.messageToJson(&m)

	log.Print(message)
}

func (l *Logger) Fatal(input error) {
	m := Message{
		Type:    "FATAL",
		Message: input.Error(),
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[FATAL] %s\n", input)
	l.messageToJson(&m)

	log.Fatal(message)
}

func (l *Logger) messageToJson(message *Message) {
	// Marshal to JSON string
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return
	}

	// Write to file
	l.logger.Println(string(jsonMessage))
}

func (l *Logger) uploadLog() {
	f, err := os.Open(logFile)
	if err != nil {
		l.Fatal(err)
	}

	// If file is empty, return
	fileInfo, err := f.Stat()
	if err != nil {
		return
	}

	if fileInfo.Size() == 0 {
		return
	}

	dataset := os.Getenv("AXIOM_DATASET")
	if dataset == "" {
		return
	}

	log.Println("Uploading log file to Axiom")

	// Rename the file to prevent loss of data
	renamedFile := fmt.Sprintf("./tmp/logs-%s.log", time.Now().Format("2006-01-02"))

	err = os.Rename(logFile, renamedFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l.logger.SetOutput(file)

	// Upload the file to a remote server
	l.uploadLogToAxiom(renamedFile, dataset)
}

func (l *Logger) uploadLogToAxiom(file string, dataset string) {
	// Upload the file to Axiom

	// 1. Open the file to ingest.
	f, err := os.Open(file)
	if err != nil {
		l.Fatal(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// 2. Wrap it in a gzip enabled reader.
	r, err := axiom.GzipEncoder(f)
	if err != nil {
		l.Fatal(err)
	}

	// 3. Initialize the Axiom API client.
	client, err := axiom.NewClient()
	if err != nil {
		l.Fatal(err)
	}

	// 4. Ingest ⚡
	// Note the JSON content type and Gzip content encoding being set because
	// the client does not auto sense them.
	res, err := client.Datasets.Ingest(context.Background(), dataset, r, axiom.NDJSON, axiom.Gzip, axiom.IngestOptions{})
	if err != nil {
		l.Fatal(err)
	}

	// 5. Make sure everything went smoothly.
	for _, fail := range res.Failures {
		l.Error(fail.Error)
	}

	// 6. Remove the file
	err = os.Remove(file)
	if err != nil {
		l.Fatal(err)
	}
}

package logging

import (
	"context"
	"fmt"
	"github.com/axiomhq/axiom-go/axiom"
	"log"
	"metar.gg/environment"
	"sync"
	"time"
)

type Logger struct {
	axiomClient *axiom.Client
	storeMutex  sync.Mutex
	eventStore  []axiom.Event
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
	client, err := axiom.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	loggerObject := &Logger{
		axiomClient: client,
	}

	// Trigger upload every 10 seconds, when axiom is configured
	if environment.Global.AxiomDataset != "" {
		go func() {
			for {
				time.Sleep(10 * time.Second)
				loggerObject.uploadLog()
			}
		}()
	}

	return loggerObject
}

func (l *Logger) Debug(input string) {
	m := Message{
		Type:    "DEBUG",
		Message: input,
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[DEBUG] %s\n", input)
	go l.messageToAxiomEvent(&m)

	log.Print(message)
}

func (l *Logger) Info(input string) {
	m := Message{
		Type:    "INFO",
		Message: input,
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[INFO] %s\n", input)
	go l.messageToAxiomEvent(&m)

	log.Print(message)
}

func (l *Logger) Warn(input string) {
	m := Message{
		Type:    "WARN",
		Message: input,
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[WARN] %s\n", input)
	go l.messageToAxiomEvent(&m)

	log.Print(message)
}

func (l *Logger) Error(input string) {
	m := Message{
		Type:    "ERROR",
		Message: input,
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[ERROR] %s\n", input)
	go l.messageToAxiomEvent(&m)

	log.Print(message)
}

func (l *Logger) Fatal(input error) {
	m := Message{
		Type:    "FATAL",
		Message: input.Error(),
		Time:    time.Now(),
	}

	message := fmt.Sprintf("[FATAL] %s\n", input)
	go l.messageToAxiomEvent(&m)

	log.Fatal(message)
}

func (l *Logger) messageToAxiomEvent(message *Message) {
	if environment.Global.AxiomDataset == "" {
		return
	}

	// Marshal to JSON string
	event := axiom.Event{
		"type":    message.Type,
		"message": message.Message,
		"time":    message.Time,
	}

	// Append to the message store
	l.storeMutex.Lock()
	l.eventStore = append(l.eventStore, event)

	// If the store is too big, upload it
	if len(l.eventStore) >= 999 {
		l.uploadLog()
	}

	l.storeMutex.Unlock()
}

func (l *Logger) uploadLog() {
	l.storeMutex.Lock()
	defer l.storeMutex.Unlock()

	storeSize := len(l.eventStore)

	if storeSize == 0 || environment.Global.AxiomDataset == "" {
		return
	}

	log.Printf("Uploading %d log events to Axiom\n", storeSize)

	// Ingest ⚡
	res, err := l.axiomClient.Datasets.IngestEvents(context.Background(), environment.Global.AxiomDataset, axiom.IngestOptions{}, l.eventStore...)
	if err != nil {
		l.Fatal(err)
	}

	// Make sure everything went smoothly.
	for _, fail := range res.Failures {
		l.Error(fail.Error)
	}

	// Clear the store
	l.eventStore = []axiom.Event{}
}

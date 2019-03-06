package api

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// LoggingData struct for structured loggin
type LoggingData struct {
	Timestamp string `json:"@timestamp,omitempty"`
	Service   string `json:"service,omitempty"`
	Thread    string `json:"thread,omitempty"`
	IP        string `json:"ip,omitempty"`
	Env       string `json:"env,omitempty"`
	Server    string `json:"server,omitempty"`

	Level        string      `json:"level,omitempty"`
	Event        string      `json:"event,omitempty"`
	Message      string      `json:"message,omitempty"`
	ID           string      `json:"Id,omitempty"`
	Raw          string      `json:"raw,omitempty"`
	RawInterface interface{} `json:"rawInterface,omitempty"`
}

// Log in json format
func Log(logD LoggingData, level, event, msg string) {
	logD.Timestamp = time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	logD.Server = hostname
	logD.Level = level
	logD.Event = event
	logD.Message = msg
	logD.Service = "Litchi-api"
	logJSON, err := json.Marshal(logD)
	if err != nil {
		log.Println("Litchi logger: Logger JSON Marshal failed !")
	}
	log.Println(string(logJSON))
}

// LogNew is to log with a new struct
func LogNew(level, event, msg string) {
	var logD LoggingData
	logD.Timestamp = time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	logD.Server = hostname
	logD.Level = level
	logD.Event = event
	logD.Message = msg
	logD.Service = "Litchi-api"
	logJSON, err := json.Marshal(logD)
	if err != nil {
		log.Println("Litchi logger: Logger JSON Marshal failed !")
	}
	log.Println(string(logJSON))
}

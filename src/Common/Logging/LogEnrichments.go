package Logging

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type LogEntry struct {
	log *log.Entry
}
type CompletedEntry LogEntry

func StartBuildingLogEntry() LogEntry {
	le := log.WithTime(time.Now())
	return LogEntry{le}
}

func (entry LogEntry) AddRequestInfo(c echo.Context) LogEntry {
	entry.log = entry.log.WithFields(log.Fields{
		"QueryString": c.QueryString(),
		"Scheme":      c.Scheme(),
		"Path":        c.Path(),
		"Method":      c.Request().Method,
	})
	return entry
}

func (entry LogEntry) AddApplicationInfo(applicationName string) LogEntry {
	hostName, _ := os.Hostname()
	exec, _ := os.Executable()
	entry.log = entry.log.WithFields(log.Fields{
		"ApplicationName": applicationName,
		"HostName":        hostName,
		"Executable":      exec,
	})
	return entry
}

func (entry LogEntry) AddErrorInfo(err interface{}) LogEntry {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(err)
	json.Unmarshal(inrec, &inInterface)
	for field, val := range inInterface {
		fmt.Println("foreach")
		entry.log = entry.log.WithField(field, val)
	}
	return entry
}

func (entry LogEntry) AddUnknownErrorInfo(err interface{}) LogEntry {
	errText, _ := json.Marshal(err)
	entry.log = entry.log.WithField("Error", string(errText))
	return entry
}

func (entry LogEntry) AsInfo() CompletedEntry {
	entry.log = entry.log.WithField("level", "info")
	ce := CompletedEntry(entry)
	return ce
}

func (entry LogEntry) AsError() CompletedEntry {
	entry.log = entry.log.WithField("level", "error")
	ce := CompletedEntry(entry)
	return ce
}
func (entry LogEntry) AsPanic() CompletedEntry {
	entry.log = entry.log.WithField("level", "panic")
	ce := CompletedEntry(entry)
	errText, _ := ce.log.String()
	fmt.Println("From AsPanic: " + errText)
	return ce
}

type Producer interface {
	Produce(message []byte, topic string)
}

func (entry CompletedEntry) WriteToKafka(producer Producer, topic string) {
	bytes, _ := entry.log.Bytes()
	producer.Produce(bytes, topic)
}

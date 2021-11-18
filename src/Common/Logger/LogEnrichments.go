package Logger

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type LogEntry struct {
	 log *log.Entry
}

func StartBuildingLogrus() LogEntry {
	le := log.WithTime(time.Now())
	return LogEntry{le}
}

func (entry LogEntry) AddRequestInfo(c echo.Context) LogEntry{
	entry.log =  entry.log.WithFields(log.Fields{
		"QueryString": c.QueryString(),
		"Scheme": c.Scheme(),
		"Path": c.Path(),
		"Method": c.Request().Method,
	})
	return entry
}

func (entry LogEntry) AddApplicationInfo(applicationName string) LogEntry{
	hostName, _ :=os.Hostname()
	exec, _ := os.Executable()
	entry.log = entry.log.WithFields(log.Fields{
		"ApplicationName": applicationName,
		"HostName": hostName,
		"Executable": exec,
	})
	return entry
}


func (entry LogEntry) AddErrorInfo(err interface{}) LogEntry{
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(err)
	json.Unmarshal(inrec, &inInterface)
	for field, val := range inInterface {
		entry.log = entry.log.WithField(field,val)
	}
	return entry
}

func (entry LogEntry) FinishBuilding() *log.Entry{
	return entry.log
}

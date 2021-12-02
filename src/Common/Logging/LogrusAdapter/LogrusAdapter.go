package LogrusAdapter

import (
	"Api.Calisma/src/Common/Logging/LogBody"
	log "github.com/sirupsen/logrus"
)

type Logrus struct {
	Body *log.Entry
}

func Create(dynBody *LogBody.DynamicBody) Logrus {
	var entry *log.Entry
	for key, value := range dynBody.Data {
		if entry == nil {
			entry = log.WithField(key, value)
		}
		entry = entry.WithField(key, value)
	}
	return Logrus{entry}
}

func (le Logrus) LogInfo() {
	le.Body.Info()
}

func (le Logrus) LogError() {
	le.Body.Error()
}

func (le Logrus) LogPanic() {
	le.Body.Panic()
}

/*LogInfo(message []byte)
GetLogMap
LogError(message []byte)
LogPanic(message []byte)*/

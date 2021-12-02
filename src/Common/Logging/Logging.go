package Logging

type Logger interface {
	Create(logBody map[string]string)
	LogInfo()
	LogError()
	LogPanic()
}

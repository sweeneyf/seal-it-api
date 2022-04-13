package logger

import "github.com/sweeneyf/seal-it-api/entity"

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger

//Service interface
type Logger interface {
	Debug(msg string, tags ...entity.KvPair)
	Info(msg string, tags ...entity.KvPair)
	Error(msg string, err error, tags ...entity.KvPair)
}

// SetLogger is the setter for log variable, it should be the only way to assign value to log
func SetLogger(newLogger Logger) {
	Log = newLogger
}

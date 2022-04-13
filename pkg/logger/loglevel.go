package logger

type LogLevel int

const (
	Info LogLevel = iota
	Debug
	Warn
	Error
)

func (ll LogLevel) String() string {
	return []string{"Info", "Debug", "Warn", "Error"}[ll]
}

// convert string to application log level , if not found default to err
func GetLogLevel(strLevel string) LogLevel {
	//default to error level
	var level LogLevel = Error
	switch strLevel {
	case "info":
		level = Info
	case "debug":
		level = Debug
	case "warn":
		level = Warn
	case "error":
		level = Error
	}
	return level
}

package agilefridge

import (
	"log"
	"net/http"
	"os"
)

type (
	Configuration struct {
		Addr string
	}
	Logger struct {
		*log.Logger
	}
)

func wrapStandardLogger(standardLogger *log.Logger) Logger {
	return Logger{
		standardLogger,
	}

}

func newStandardLogger(prefix string) *log.Logger {
	return log.New(
		os.Stderr,
		prefix,
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile,
	)
}

func NewLogger() Logger {
	return wrapStandardLogger(newStandardLogger(""))
}

func (logger *Logger) WithRequestID(requestID string) Logger {
	return wrapStandardLogger(newStandardLogger(requestID))
}

func getFromEnvWithDefault(path string, fallback string) string {
	if value, ok := os.LookupEnv(path); ok {
		return value
	} else {
		return fallback
	}
}

func GetConfigurationFromEnvironment() Configuration {
	return Configuration{
		Addr: getFromEnvWithDefault("ADDR", ":3000"),
	}
}

func Run(configuration Configuration) int {
	logger := NewLogger()

	server := &http.Server{
		Addr:    configuration.Addr,
		Handler: NewRouter(&logger),
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Println(err)
		return 1
	}

	return 0
}

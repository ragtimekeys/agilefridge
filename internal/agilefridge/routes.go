package agilefridge

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

func NewRouter(logger *Logger) *http.ServeMux {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		requestIDBytes := [24]byte{}
		_, err := rand.Read(requestIDBytes[:])
		if err != nil {
			res.WriteHeader(500)
			res.Write([]byte("Cannot generate a request ID"))
			return
		}

		requestID := base64.StdEncoding.EncodeToString(requestIDBytes[:])

		requestLogger := logger.WithRequestID(requestID + " ")
		requestLogger.Println("Hello world")
	})

	return serveMux
}

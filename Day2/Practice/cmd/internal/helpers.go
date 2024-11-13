package internal

import (
	"net/http"
)

// Дополнительные типы для работы пакета internal

type Message struct {
	Message string `json:"message"`
}

func WriteHeaderAndSaveStatus(statusCode int, w http.ResponseWriter) int {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	return statusCode
}

func WithLoggingHandle(h http.Handler) func(w http.ResponseWriter, r *http.Request) {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		GlobalSugar.Infoln(
			"statusCode", StatusCode,
		)
	}
	return logFn
}

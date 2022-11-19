package logger

import (
	"net/http"
)

type MockLogger struct{}

func NewMockLogger() MockLogger {
	return MockLogger{}
}

func (MockLogger) Info(message string) {}

func (MockLogger) Error(message string) {}

func (MockLogger) RequestLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

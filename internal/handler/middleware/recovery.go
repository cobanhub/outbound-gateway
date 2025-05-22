package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// This middleware recovers from panics and writes a 500 Internal Server Error response.
func (m *Middleware) RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(m.timeout))
		defer cancel()

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				log.Printf("Request timed out: %v", ctx.Err())
				http.Error(w, "Request Timeout", http.StatusRequestTimeout)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// CorrelationIDMiddleware adds a correlation ID to the request and response headers.
func CorrelationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.NewString()
		}

		w.Header().Set("X-Correlation-ID", correlationID)

		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		log.Printf("[%s] %s %s %v", correlationID, r.Method, r.URL.Path, duration)
	})
}

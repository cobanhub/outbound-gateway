package middleware

import (
	"net/http"
	"time"
)

type (
	Middleware struct {
		timeout time.Duration
	}

	MiddlewareOptions struct {
		Timeout time.Duration
	}

	MiddlewareInterface interface {
		RecoveryMiddleware(next http.Handler) http.Handler
		CorrelationIDMiddleware(next http.Handler) http.Handler
	}
)

// MiddlewareOption is a function that applies a configuration option to the Middleware.
func NewMiddleware(opts MiddlewareOptions) *Middleware {
	return &Middleware{
		timeout: opts.Timeout,
	}
}

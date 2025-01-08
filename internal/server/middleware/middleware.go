package middleware

import (
	"hot-coffee/pkg/logger"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	body        string
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) Body() string {
	return rw.body
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}

	rw.body = string(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func LoggingMiddleware(logger *logger.CustomLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			logger.InfoLogger.Printf("Started %s %s", r.Method, r.URL.Path)

			wrapped := wrapResponseWriter(w)

			next.ServeHTTP(wrapped, r)

			if wrapped.Status() >= 400 {
				logger.ErrorLogger.Printf("%s %s %d %s", r.Method, r.URL.Path, wrapped.Status(), wrapped.Body())
			}

			if wrapped.Status() < 400 {
				logger.InfoLogger.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
			}
		})
	}
}

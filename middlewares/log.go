package middleware

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func Log(next httprouter.Handle) httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Start Time
		startTime := time.Now()

		// Status Writer
		sw := statusWriter{ResponseWriter: w}

		// Next
		if next != nil {
			next(&sw, r, ps)
		}

		// End Time
		endTime := time.Now()

		// Request Duration
		duration := endTime.Sub(startTime).String()

		// Log Request End
		zap.L().Info("Request",
			zap.String("Method", r.Method),
			zap.String("Path", r.URL.Path),
			zap.String("Query", r.URL.RawQuery),
			zap.Time("Timestamp", endTime),
			zap.String("Duration", duration),
			zap.Int("Status", sw.status),
		)
	}
	return fn
}

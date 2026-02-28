package admin

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

// putting a counter on the /app requests
// then assembling tools to deal with the counter
type APIConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *APIConfig) MiddlewareMetricsIncrement(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
// /metrics
func (cfg *APIConfig) HandleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	hits := cfg.fileserverHits.Load()
	page := fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, hits)
	w.Write([]byte(page))
}
// /reset
func (cfg *APIConfig) HandleReset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	cfg.fileserverHits.Store(0)
	page := fmt.Sprintf("Hits reset to %d", cfg.fileserverHits.Load())
	w.Write([]byte(page))
}
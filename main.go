package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
//	"sync"
	"sync/atomic"
//	"time"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handleMetricEndpoint(w http.ResponseWriter, r *http.Request) {
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

func (cfg *apiConfig) handleResetEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	cfg.fileserverHits.Store(0)
	page := fmt.Sprintf("Hits reset to %d", cfg.fileserverHits.Load())
	w.Write([]byte(page))
}

func handleHealthEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	const page = `OK`
	w.Write([]byte(page))
}

func main() {
	serveMux := http.NewServeMux()
	apiCfg := apiConfig{

	}

	serverRoot := os.Getenv("GOSERVER_ROOT")
	serveMux.Handle(
		"/app/",
		apiCfg.middlewareMetricsInc(
			http.StripPrefix(
				"/app/",
				http.FileServer(
					http.Dir(serverRoot),
				),
			),
		),
	)
	serveMux.HandleFunc(
		"GET /api/healthz",
		handleHealthEndpoint,
	)
	serveMux.HandleFunc(
		"GET /admin/metrics",
		apiCfg.handleMetricEndpoint,
	)
	serveMux.HandleFunc(
		"POST /admin/reset",
		apiCfg.handleResetEndpoint,
	)

	port := os.Getenv("GOSERVER_PORT")
	srv := http.Server{
		Handler: serveMux,
		Addr: ":" + port,
		// WriteTimeout: 30 * time.Second,
		// ReadTimeout: 30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Printf("server started on %s\n", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

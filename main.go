package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
//	"sync"
	"sync/atomic"
//	"time"
)

// api API
// /healthz
func handleHealthEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	const page = `OK`
	w.Write([]byte(page))
}
// /validate_chirp
func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{Error: msg,})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: unable to marshal JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	// made it here so we haven't had a server error yet
	// if it fits it ships
	w.WriteHeader(code)
	w.Write(dat)
}

// ensure chirp is 140 characters or less
func handleValidateChirpEndpoint(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid bool `json:"valid"`
	}

	// are we dealing with a chirp?
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to decode request body", err)
		return
	}
	// we have a chirp,	is it too long?
	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
			respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
			return
	}
	
	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
	})
}

// putting a counter on the /app requests
// then assembling tools to deal with the counter
type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
// /metrics
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
// /reset
func (cfg *apiConfig) handleResetEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	cfg.fileserverHits.Store(0)
	page := fmt.Sprintf("Hits reset to %d", cfg.fileserverHits.Load())
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
		"POST /api/validate_chirp",
		handleValidateChirpEndpoint,
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

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
type responseBody interface {
	doJSONMarshal() ([]byte, error)
}

type errorResponseBody struct {
	body string `json:"error"`
}

func (e errorResponseBody) doJSONMarshal() ([]byte, error) {
	return json.Marshal(e)
}

type successResponseBody struct {
	body bool `json:"valid"`
}

func (s successResponseBody) doJSONMarshal() ([]byte, error) {
	return json.Marshal(s)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respBody := errorResponseBody{
		body: msg,
	}
	respondWithJSON(w, code, respBody)
}

func respondWithJSON(w http.ResponseWriter, code int, payload responseBody) {
	dat, err := payload.doJSONMarshal()
	if err != nil {
		log.Printf("Error: unable to marshal JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	// made it here so we haven't had a server error yet
	// if it fits it ships
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
// ensure chirp is 140 characters or less
func handleValidateChirpEndpoint(w http.ResponseWriter, r *http.Request) {
	const maxChirpLength = 140
	var decoded struct {
		body string `json:"body"`
	}
	// are we dealing with a chirp?
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&decoded)
	if err != nil {
		respondWithError(w, 400, "unable to decode request body")
	} else {
		// we have a chirp,	is it too long?
		if len(decoded.body) > maxChirpLength {
			respondWithError(w, 400, "Chirp is too long")
	  } else {
	  	respondWithJSON(w, 200, successResponseBody{body: true,})
	  }
	}
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

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
//	"time"
)

func main() {
	m := http.NewServeMux()

	serverRoot := os.Getenv("GOSERVER_ROOT")
	m.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir(serverRoot))))
	m.HandleFunc("/healthz", handleHealthEndpoint)

	port := os.Getenv("GOSERVER_PORT")
	srv := http.Server{
		Handler:      m,
		Addr:         ":" + port,
		// WriteTimeout: 30 * time.Second,
		// ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Printf("server started on %s\n", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handleHealthEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	const page = `OK`
	w.Write([]byte(page))
}
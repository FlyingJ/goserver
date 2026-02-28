package main

import (
	"log"
	"net/http"
	"os"
	"github.com/FlyingJ/goserver/internal/admin"
	"github.com/FlyingJ/goserver/internal/api"
)

func main() {
	serveMux := http.NewServeMux()
	apiCfg := admin.APIConfig{}

	serverRoot := os.Getenv("GOSERVER_ROOT")
	serveMux.Handle(
		"/app/",
		apiCfg.MiddlewareMetricsIncrement(
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
		api.HandleHealth,
	)
	serveMux.HandleFunc(
		"POST /api/validate_chirp",
		api.HandleValidateChirp,
	)
	serveMux.HandleFunc(
		"GET /admin/metrics",
		apiCfg.HandleMetrics,
	)
	serveMux.HandleFunc(
		"POST /admin/reset",
		apiCfg.HandleReset,
	)

	port := os.Getenv("GOSERVER_PORT")
	srv := http.Server{
		Handler: serveMux,
		Addr: ":" + port,
	}

	// this blocks forever, until the server has an unrecoverable error
	log.Printf("server started on %s\n", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
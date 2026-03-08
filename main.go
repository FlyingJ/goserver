package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/FlyingJ/goserver/internal/admin"
	"github.com/FlyingJ/goserver/internal/api"
	"github.com/FlyingJ/goserver/internal/database"
)

func main() {
	apiCfg := admin.APIConfig{}

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	serverRoot := os.Getenv("GOSERVER_ROOT")

	db, err := sql.Open("postgres", dbURL)
	apiCfg.DBQueries = database.New(db)

	serveMux := http.NewServeMux()
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
		"POST /api/users",
		api.HandleUsers,
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
	err = srv.ListenAndServe()
	log.Fatal(err)
}
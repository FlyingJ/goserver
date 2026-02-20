package main

import (
	"fmt"
	"log"
	"net/http"
//	"os"
//	"time"
)

func main() {
	m := http.NewServeMux()

	serverRoot := "./public"
	m.Handle("/", http.FileServer(http.Dir(serverRoot)))

	// port := os.Getenv("PORT")
	port := ":8080"
	srv := http.Server{
		Handler:      m,
		Addr:         port,
		// WriteTimeout: 30 * time.Second,
		// ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Printf("server started on %s\n", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

/*
func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p>Welcome to Chirpy</p>
</body>
</html>
`
	w.Write([]byte(page))
}
*/
package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// /validate_chirp
func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errorResponse{Error: msg,})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
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
func HandleValidateChirp(w http.ResponseWriter, r *http.Request) {
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
		RespondWithError(w, http.StatusInternalServerError, "unable to decode request body", err)
		return
	}
	if params == (parameters{}) {
		RespondWithError(w, http.StatusBadRequest, "empty payload", nil)
		return
	}
	// we have a chirp,	is it too long?
	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
			RespondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
			return
	}

	RespondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
	})
}
package api

import (
	"encoding/json"
	"github.com/FlyingJ/goserver/internal/constants"
	"github.com/FlyingJ/goserver/internal/util"
	"net/http"
)

// /validate_chirp
func HandleValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	// are we dealing with a chirp?
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "unable to decode request body", err)
		return
	}
	if params == (parameters{}) {
		util.RespondWithError(w, http.StatusBadRequest, "empty payload", nil)
		return
	}

	// we have a chirp,	is it too long?
	chirp := params.Body
	if len(chirp) > constants.MaxChirpLength {
			util.RespondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
			return
	}

	// apply profanity filter and send result
	util.RespondWithJSON(w, http.StatusOK, returnVals{CleanedBody: util.Censor(chirp),})
}
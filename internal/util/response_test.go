package util_test

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlyingJ/goserver/internal/util"
)


// func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	dat, err := json.Marshal(payload)
// 	if err != nil {
// 		log.Printf("Error: unable to marshal JSON: %s", err)
// 		w.WriteHeader(500)
// 		return
// 	}
// 	// made it here so we haven't had a server error yet
// 	// if it fits it ships
// 	w.WriteHeader(code)
// 	w.Write(dat)
// }
func TestRespondWithJSON(t *testing.T) {
	w := httptest.NewRecorder()
	code := http.StatusOK
	type respVals struct {
		body string
	}
	util.RespondWithJSON(w, code, respVals{body:"test",})
	
	if w.Code != http.StatusOK {
		t.Errorf("RespondWithJSON response is not OK")
	} else {
		log.Println("RespondWithJSON response is OK")
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("RespondWithJSON response Content-Type sucks: %s", contentType)
	} else {
		log.Printf("RespondWithJSON response Content-Type is good: %s", contentType)
	}

	decoder := json.NewDecoder(w.Body)
	var r respVals
	err := decoder.Decode(&r)
	if err != nil {
		t.Errorf("unable to decode response from RespondWithJSON")
	}
	if r.body != "test" {
	    t.Errorf("RespondWithJSON response body not OK")
	} else {
		log.Printf("RespondWithJSON response is good: %s", r.body)
	}
}

// func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
//     if err != nil {
//         log.Println(err)
//     }
//     if code > 499 {
//         log.Printf("Responding with 5XX error: %s", msg)
//     }
//     type errorResponse struct {
//         Error string `json:"error"`
//     }
//         RespondWithJSON(w, code, errorResponse{Error: msg,})
// }
func TestRespondWithError(t *testing.T) {
	return
}


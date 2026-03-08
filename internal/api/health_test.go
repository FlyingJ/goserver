package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlyingJ/goserver/internal/api"
)

/*
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	const page = `OK`
	w.Write([]byte(page))
}
*/
func TestHandleHealth(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "test", nil)
	w := httptest.ResponseRecorder

	expectation = struct{
		Code: 200,
		HeaderMap: map[string]string{
			"Content-Type": "text/plain; charset=utf-8",
		},
		Body: []byte(`OK`)
	}

	req := httptest.NewRequest("GET", "test", nil)
	w := httptest.NewRecorder()
	api.HandleHealth(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	
	if w.Code != expectation.Code {
		t.Errorf("HandleHealth response is not OK")
	}

	contentType := w.HeaderMap().Get("Content-Type")
	if contentType != expectation.HeaderMap["Content-Type"] {
		t.Errorf("HandleHealth response Content-Type sucks: %s", contentType)
	}

	if w.Body != expectation.Body {
		t.Errorf("HandleHealth response body not OK")
	}
}




//import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// )

// func main() {
// 	handler := func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "<html><body>Hello World!</body></html>")
// 	}
//
//	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
// 	w := httptest.NewRecorder()
// 	handler(w, req)

// 	resp := w.Result()
// 	body, _ := io.ReadAll(resp.Body)

// 	fmt.Println(resp.StatusCode)
// 	fmt.Println(resp.Header.Get("Content-Type"))
// 	fmt.Println(string(body))

// }
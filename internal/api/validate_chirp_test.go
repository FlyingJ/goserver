package api_test

import (
	"net/http"
	"testing"

	"github.com/FlyingJ/goserver/internal/api"
)

func TestValidateChirp(t *testing.T) {
	cases := []struct{
		input       *http.Request
		expectation *http.Response
	}{
		{
			input: "hello world",
			expectation: "hello world",
		},
		{
			input: "hello you silly fornax",
			expectation: "hello you silly ****",
		},
		{
			input: "",
			expectation: "",
		},
		{
			input: "Fornax!",
			expectation: "Fornax!",
		},
	}
	for _, c := range cases {
		in := c.input
		xpct := c.expectation
		w := 
		res, err := api.HandleValidateChirp(w, input)
		if result != expectation {
			t.Errorf(
				"Error: %s and %s do not match",
				result,
				expectation,
			)
		}
	}
}
}
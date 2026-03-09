package util_test

import (
	"testing"

	"github.com/FlyingJ/goserver/internal/util"
)

func TestCensor(t *testing.T) {
	cases := []struct{
		input       string
		expectation string
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
		input := c.input
		expectation := c.expectation
		result := util.Censor(input)
		if result != expectation {
			t.Errorf(
				"Error: %s and %s do not match",
				result,
				expectation,
			)
		}
	}
}
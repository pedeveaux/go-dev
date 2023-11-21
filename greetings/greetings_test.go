package greetings

import (
	"testing"
)

func TestHello(t *testing.T) {
	tt := []struct {
		in  string
		exp string
	}{
		{
			in:  "Mike",
			exp: "Hi, Mike. Welcome to the Go show !"},
		{
			in:  "George",
			exp: "Hi, George. Welcome to the Go show !"},
	}
	for _, tc := range tt {
		if actual, err := Hello(tc.in); actual != tc.exp {
			t.Errorf("Expected Hello to be %s but it was %s. Error: %s", tc.exp, actual, err.Error())
		}
	}
}

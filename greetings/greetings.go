package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {
	// If not name was given return an error
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a name was received, return a value that embeds the name
	// in a greeting message
	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome to the Go show !", name)
	return message, nil
}

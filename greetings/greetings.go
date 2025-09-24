package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name) // declare and initialize a variable at the same time
	return message, nil
}

func randomFormat() string {
	formats := []string {
		"hi, %v. welcome!",
		"great to see you, %v!",
		"hail, %v! well met!",
	}

	return formats[rand.Intn(len(formats))]
}
package greetings

import (
	"testing"
	"regexp"
)

// Hello should return a personalized greeting 
func TestHelloName(t *testing.T) {
	name := "Gladys"

	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Hello should return error if no name is provided
func TestHelloEmpty (t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
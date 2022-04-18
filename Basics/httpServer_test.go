package basics

import "testing"

func RunServer_and_call_API() {
	go RunServer()
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg := Hello("Mario")
	if msg != "Hello Mario!" {
		t.Fatalf(`Hello("") = %q, want "", error`, msg)
	}
}

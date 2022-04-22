package basics

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestRunServer_and_call_API(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	server := &http.Server{
		Addr: ":9001",
	}

	go CreateServer(server, wg)
	defer server.Shutdown(context.Background())

	// wait until server is up and running
	wg.Wait()
	time.Sleep(time.Millisecond * 1000)

	req, err := http.NewRequest("GET", "http://localhost:9001", nil)
	if err != nil {
		log.Fatal()
	}

	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}
	// Defer the closing of the body
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	str := string(b)
	if err != nil {
		t.Fatalf("Failed to read response body!")
	}

	if str != "Hello Mario!" {
		t.Fatalf(`Hello("") = %q, want "", error`, str)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg := Hello("Mario")

	if msg != "Hello Mario!" {
		t.Fatalf(`Hello("") = %q, want "", error`, msg)
	}
}

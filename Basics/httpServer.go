/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type GreetingRequest struct {
	Name string `json:"name"`
}

type GreetingResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	name := "Mario"

	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Has("name") {
			name = r.URL.Query()["name"][0]
		}

		io.WriteString(w, fmt.Sprintf("Hello %v!", name))
	case http.MethodPost:
		// https://medium.com/what-i-talk-about-when-i-talk-about-technology/go-code-snippet-json-encoder-and-json-decoder-818f81864614

		var req GreetingRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		name = req.Name

		resp := GreetingResponse{
			Name:    name,
			Message: fmt.Sprintf("Hello %v!", name),
		}

		err = json.NewEncoder(w).Encode(resp) // note: you cannot assign ':="" err but '=' works
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

}

// Create a server on specific port. WaitGroup is notified when server is initialized and can be nil.
func CreateServer(server *http.Server, wg *sync.WaitGroup) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handle)

	server.Handler = mux

	if wg != nil {
		wg.Done()
	}

	err := server.ListenAndServe()

	if err != nil && !strings.Contains(err.Error(), "Server closed") {
		log.Fatal(err)
	}
}

func main() {
	const port int = 9000
	fmt.Printf("starting server on port %v ...\n", port)

	serverAddr := fmt.Sprintf(":%v", port)

	server := &http.Server{
		Addr: serverAddr,
	}
	CreateServer(server, nil)

	time.AfterFunc(time.Second*10, func() {
		server.Shutdown(context.Background())
	})

	fmt.Println("shutting down server after 10s ...")
}

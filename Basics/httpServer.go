/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package basics

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func Hello(text string) string {
	return fmt.Sprintf("Hello %v!", text)
}

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Mario!")
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

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	server := &http.Server{
		Addr: ":9000",
	}
	CreateServer(server, nil)

	time.AfterFunc(time.Second*10, func() {
		server.Shutdown(context.Background())
	})
}

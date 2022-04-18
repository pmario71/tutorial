/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package basics

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Hello(text string) string {
	return fmt.Sprintf("Hello %v!", text)
}

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Mario!")
}

func RunServer() {
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	err := http.ListenAndServe(":"+portNumber, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	RunServer()
}

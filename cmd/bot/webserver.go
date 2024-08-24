package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Discord bot ok")
}

func runServer(wg *sync.WaitGroup) {
	defer wg.Done()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //
	}

	http.HandleFunc("/", homeHandler)
	fmt.Printf("Starting server on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

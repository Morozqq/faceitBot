package main

import (
	"fmt"
	"net/http"
	"sync"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Discord bot ok")
}

func runServer(wg *sync.WaitGroup) {
	defer wg.Done()
	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

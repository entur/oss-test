package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world\n"))
}

func main() {
	http.HandleFunc("/actuator/health/", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

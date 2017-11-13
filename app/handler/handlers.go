package handler

import (
	"fmt"
	"net/http"
	"os"
)

// Main Handler
func Main(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, World!\nHostname: %s\n", hostname)
}

// Healthz Handler for monitoring app health
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Chilling\n")
}

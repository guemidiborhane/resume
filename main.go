package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:embed index.html
var html []byte

func main() {
	// Check if this is a healthcheck call
	if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
		healthcheck()
		return
	}

	port := getPort()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=3600")
		w.Write(html)
	})

	mux.HandleFunc("GET /_healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	})

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server starting on http://localhost%s\n", addr)
	fmt.Printf("Health check: http://localhost%s/_healthz\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func healthcheck() {
	port := getPort()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/_healthz", port))
	if err != nil {
		fmt.Printf("Health check failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Health check failed with status: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	fmt.Println("Health check passed")
	os.Exit(0)
}

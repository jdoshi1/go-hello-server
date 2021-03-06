package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

const (
	ServerPort = "8080"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Guest"
	}

	log.Printf("Endpoint: %s\n", r.URL.Path)
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s!\nCurrent time: %s\n", name, time.Now().String())))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint: %s\n", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Server is healthy!\n")))

}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint: %s\n", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Server is ready!\n")))
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/readiness", readinessHandler)

	srv := &http.Server{
		Handler: r,
		Addr: fmt.Sprintf(":%s", ServerPort),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal
	<-interruptChan

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
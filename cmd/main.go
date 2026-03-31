package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	port := ":8080"

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	})

	srv := &http.Server{
		Addr:         port,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}

	go func() {
		log.Printf("Started server on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error occured while starting server: %v", err)
		}
	}()

	<-sigs
	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

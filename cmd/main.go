package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahmed-cmyk/drawsync/internal/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup DB
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Failed to initialize driver: %v\n", err)
	}
	defer db.Close()

	// Actually verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Could not connect to DB: %v\n", err)
	}

	h := handlers.New(db)

	// Initialise Tables
	if err := h.InitialiseTables(); err != nil {
		log.Fatalf("Failed to initialise tables: %v\n", err)
	}

	mux := http.NewServeMux()
	port := ":8080"

	mux.HandleFunc("/room/{name}", h.CreateRoom)
	mux.HandleFunc("/ws", h.UpgradeConnection)

	srv := &http.Server{
		Addr:         port,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Started server on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error occured while starting server: %v\n", err)
		}
	}()

	<-sigs
	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exiting")
}

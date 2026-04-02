package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type Handler struct {
	db *sql.DB
}

func New(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) UpgradeConnection(w http.ResponseWriter, r *http.Request) {
	// 1. Upgrade the connection
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade: %v", err)
		return
	}
	// Ensure cleanup happens ONLY when the loop below finishes
	defer c.CloseNow()

	log.Println("User connected to DrawSync")

	// 2. Keep the handler alive with a loop
	for {
		var v any
		// Use r.Context() here because if the user closes their browser,
		// the context cancels and this loop breaks automatically.
		err = wsjson.Read(r.Context(), c, &v)
		if err != nil {
			log.Printf("Connection closed or error: %v", err)
			break // Exit the loop to trigger the defers and clean up
		}

		log.Printf("Received drawing data: %v", v)

		// This is where you'd eventually call h.Hub.Broadcast(v)
	}

	log.Println("User disconnected")
	c.Close(websocket.StatusNormalClosure, "Session ended")
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	_, err := h.db.Exec("INSERT INTO rooms (name) VALUES (?)", name)
	if err != nil {
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		log.Printf("Failed to create room: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created room successfully\n"))
}

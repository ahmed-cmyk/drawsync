package handlers

import (
	"database/sql"
	"log"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func New(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
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

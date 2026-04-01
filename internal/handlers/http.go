package handlers

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Write(fmt.Appendf([]byte{}, "You created a room for %s\n", name))
}

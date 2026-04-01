package handlers

func (h *Handler) InitialiseTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS rooms (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
	`

	_, err := h.db.Exec(query)
	return err
}

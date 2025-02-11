package server

import (
	"net/http"
)

type Note struct {
	ID        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Note      string `json:"note"`
}

func (cfg *Config) HandlerGetNotes(w http.ResponseWriter, r *http.Request) {

	dbNotes, err := cfg.Db.GetAllNotes(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error getting notes", err)
		return
	}

	var notes []Note
	for _, dbNote := range dbNotes {
		notes = append(notes, Note{
			ID:        dbNote.ID,
			CreatedAt: dbNote.CreatedAt,
			UpdatedAt: dbNote.UpdatedAt,
			Note:      dbNote.Note,
		},
		)
	}

	respondWithJSON(w, http.StatusOK, notes)

}

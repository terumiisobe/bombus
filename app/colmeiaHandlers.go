package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/terumiisobe/bombus/service"
)

type Colmeia struct {
	ID           int     `json:"id"`
	ColmeiaID    int     `json:"colmeia_id"` // Additional visual ID
	QRCode       *string `json:"qr_code"`    // Can be NULL
	Species      string  `json:"species"`
	StartingDate string  `json:"starting_date"`
	Status       string  `json:"status"`
}

type ColmeiaHandler struct {
	s service.ColmeiaService
}

func (ch *ColmeiaHandler) getAllColmeias(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	species := r.URL.Query().Get("species")

	colmeias, _ := ch.s.GetAllColmeia(status, species)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(colmeias)
}

func (ch *ColmeiaHandler) getColmeia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	colmeias, err := ch.s.GetColmeia(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, colmeias)
	}
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

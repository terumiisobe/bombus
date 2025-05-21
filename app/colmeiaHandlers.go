package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"bombus/domain"
	"bombus/service"
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

	colmeias, err := ch.s.GetAllColmeia(status, species)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, colmeias)
	}
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

func (ch *ColmeiaHandler) createColmeia(w http.ResponseWriter, r *http.Request) {
	var colmeia domain.Colmeia
	if err := json.NewDecoder(r.Body).Decode(&colmeia); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	err := ch.s.CreateColmeia(colmeia)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, nil)
	}
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

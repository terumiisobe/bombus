package app

import (
	"encoding/json"
	"fmt"
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
	//mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)
	//colmeias := []Colmeia{
	//	{123, 123, nil, 1, mockTime, 1},
	//	{456, 456, nil, 2, mockTime, 2},
	//}
	colmeias, _ := ch.s.GetAllColmeia()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(colmeias)
}

func (ch *ColmeiaHandler) getColmeia(w http.ResponseWriter, r *http.Request) {
	//mockTime := time.Date(2025, time.April, 15, 10, 30, 0, 0, time.UTC)
	//colmeias := []Colmeia{
	//	{123, 123, nil, 1, mockTime, 1},
	//	{456, 456, nil, 2, mockTime, 2},
	//}
	vars := mux.Vars(r)
	id := vars["id"]

	colmeias, err := ch.s.GetColmeia(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(colmeias)

	}
}

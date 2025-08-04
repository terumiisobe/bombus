package app

import (
	"encoding/json"
	"log"
	"net/http"

	"bombus/domain"
	"bombus/errs"
	"bombus/mapper"
	"bombus/service"

	"github.com/gorilla/mux"
)

type ColmeiaHandler struct {
	s service.ColmeiaService
}

func (ch *ColmeiaHandler) getAllColmeias(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	species := r.URL.Query().Get("species")

	colmeias, err := ch.s.GetAllColmeia(status, species)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, mapper.ToDTOList(colmeias))
}

func (ch *ColmeiaHandler) getColmeia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	colmeia, err := ch.s.GetColmeia(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, mapper.ToDTO(*colmeia))
}

func (ch *ColmeiaHandler) createColmeia(w http.ResponseWriter, r *http.Request) {
	var colmeia domain.Colmeia
	if err := json.NewDecoder(r.Body).Decode(&colmeia); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := ch.s.CreateColmeia(colmeia)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, nil)
}

func (ch *ColmeiaHandler) countBySpecies(w http.ResponseWriter, r *http.Request) {
	countBySpecies, err := ch.s.CountBySpecies()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, mapper.SpeciesCountToString(countBySpecies))
}

func (ch *ColmeiaHandler) countBySpeciesAndStatus(w http.ResponseWriter, r *http.Request) {
	countBySpeciesAndStatus, err := ch.s.CountBySpeciesAndStatus()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, mapper.SpeciesAndStatusCountToString(countBySpeciesAndStatus))
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		jsonErr := errs.NewJsonConversionError(err.Error())
		log.Printf(jsonErr.Message)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(jsonErr.Code)

		errorData, _ := json.Marshal(jsonErr.AsMessage())
		w.Write(errorData)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}

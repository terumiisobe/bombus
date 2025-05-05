package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/terumiisobe/bombus/domain"
	"github.com/terumiisobe/bombus/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	handler := ColmeiaHandler{service.NewColmeiaService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/colmeia", handler.getAllColmeias).Methods(http.MethodGet)
	router.HandleFunc("/greet", greet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

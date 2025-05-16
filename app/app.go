package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/terumiisobe/bombus/config"
	"github.com/terumiisobe/bombus/domain"
	"github.com/terumiisobe/bombus/service"
)

var AppConfig *config.Config

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = config.NewConfig()

	router := mux.NewRouter()

	// wiring
	colmeiaHandler := ColmeiaHandler{service.NewColmeiaService(domain.NewColmeiaRepositoryDB())}
	WhatsappHandler := WhatsappHandler{}

	// define routes
	router.HandleFunc("/colmeias", colmeiaHandler.getAllColmeias).Methods(http.MethodGet)
	router.HandleFunc("/colmeias/{id:[0-9]+}", colmeiaHandler.getColmeia).Methods(http.MethodGet)
	router.HandleFunc("/colmeias", colmeiaHandler.createColmeia).Methods(http.MethodPost)

	router.HandleFunc("/webhook", WhatsappHandler.webhookHandler).Methods(http.MethodPost)

	router.HandleFunc("/greet", greet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

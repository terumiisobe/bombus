package app

import (
	"fmt"
	"log"
	"net/http"

	"bombus/config"
	"bombus/domain"
	"bombus/repository"
	"bombus/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	// TODO: replace to actual and not stubs
	colmeiaService := service.NewColmeiaService(repository.NewColmeiaRepositoryStub())
	colmeiaHandler := ColmeiaHandler{colmeiaService}
	chatbotHandler := ChatbotHandler{service.NewChatbotService(domain.NewInteractionRepositoryStub(), colmeiaService)}

	// define routes
	router.HandleFunc("/colmeias", colmeiaHandler.getAllColmeias).Methods(http.MethodGet)
	router.HandleFunc("/colmeias/{id:[0-9]+}", colmeiaHandler.getColmeia).Methods(http.MethodGet)
	router.HandleFunc("/colmeias", colmeiaHandler.createColmeia).Methods(http.MethodPost)

	router.HandleFunc("/webhook", chatbotHandler.handle).Methods(http.MethodPost)

	router.HandleFunc("/greet", greet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

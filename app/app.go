package app

import (
	"bombus/config"
	"bombus/domain"
	"bombus/repository"
	"bombus/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var AppConfig *config.Config

func Start() {

	env := os.Getenv("ENV")
	fmt.Println("ENV value is:", env) // Debugging

	if os.Getenv("ENV") != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file (non-fatal):", err)
		}
	}

	AppConfig = config.NewConfig()

	router := mux.NewRouter()

	// wiring
	// TODO: replace to actual and not stubs
	colmeiaService := service.NewColmeiaServiceImplDefault(repository.NewColmeiaRepositoryImplStub())
	colmeiaHandler := ColmeiaHandler{colmeiaService}
	//TODO: replace to chatbot AI
	chatbotHandler := ChatbotHandler{service.NewChatbotService(domain.NewInteractionRepositoryStub(), colmeiaService)}

	// define routes
	router.HandleFunc("/colmeias", colmeiaHandler.getAllColmeias).Methods(http.MethodGet)
	router.HandleFunc("/colmeias/{id:[0-9]+}", colmeiaHandler.getColmeia).Methods(http.MethodGet)
	router.HandleFunc("/colmeias", colmeiaHandler.createColmeia).Methods(http.MethodPost)
	router.HandleFunc("/colmeias/count-by-species", colmeiaHandler.countBySpecies).Methods(http.MethodGet)
	router.HandleFunc("/colmeias/count-by-species-and-status", colmeiaHandler.countBySpeciesAndStatus).Methods(http.MethodGet)

	router.HandleFunc("/webhook", chatbotHandler.handle).Methods(http.MethodPost)

	router.HandleFunc("/greet", greet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

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
	"time"

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
	colmeiaService := service.NewColmeiaServiceImplDefault(repository.NewColmeiaRepositoryImplStubCustomData(generateCustomData()))
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(router)))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func intPtr(i int) *int {
	return &i
}

// TODO: remove this function when we have the real data
func generateCustomData() []domain.Colmeia {
	mockTime := time.Date(2025, time.August, 12, 00, 00, 0, 0, time.UTC)
	colmeias := []domain.Colmeia{}

	codesForMC := []int{1, 2, 6, 7, 8, 9, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 43, 45, 72, 86, 87, 91, 697}
	codesForGU := []int{3, 10, 70}
	codesForGA := []int{11}
	codesForJT := []int{44, 69, 71, 73, 74, 75, 77, 78, 79, 80, 83, 85}
	codesForMD := []int{76, 88}
	codesForCN := []int{81}
	codesForEM := []int{4, 5, 84}
	codesForTB := []int{90, 92, 93}

	for _, code := range codesForMC {
		colmeia := domain.Colmeia{
			ID:           code,
			ColmeiaID:    intPtr(code),
			Species:      domain.MeliponaQuadrifasciata,
			Status:       domain.Developing,
			StartingDate: mockTime,
		}
		colmeias = append(colmeias, colmeia)
	}

	for _, code := range codesForGU {
		colmeia := domain.Colmeia{
			ID:           code,
			ColmeiaID:    intPtr(code),
			Species:      domain.PlebeiaGigantea,
			Status:       domain.Developing,
			StartingDate: mockTime,
		}
		colmeias = append(colmeias, colmeia)
	}

	for _, code := range codesForGA {
		colmeia := domain.Colmeia{
			ID:           code,
			ColmeiaID:    intPtr(code),
			Species:      domain.MeliponaBicolor,
			Status:       domain.Developing,
			StartingDate: mockTime,
		}
		colmeias = append(colmeias, colmeia)
	}

	for _, code := range codesForJT {
		colmeia := domain.Colmeia{
			ID:           code,
			ColmeiaID:    intPtr(code),
			Species:      domain.TetragosniscaAngustula,
			Status:       domain.Developing,
			StartingDate: mockTime,
		}
		colmeias = append(colmeias, colmeia)
	}

	for _, code := range codesForMD {
		colmeia := domain.Colmeia{
			ID:           code,
			ColmeiaID:    intPtr(code),
			Species:      domain.MeliponaMarginata,
			Status:       domain.Developing,
			StartingDate: mockTime,
		}
		colmeias = append(colmeias, colmeia)
	}

	for _, code := range codesForCN {
		colmeia := domain.Colmeia{
			ID:           code,
			ColmeiaID:    intPtr(code),
			Species:      domain.ScaptotrigonaDepilis,
			Status:       domain.Developing,
			StartingDate: mockTime,
		}
		colmeias = append(colmeias, colmeia)
	}

	for _, code := range codesForEM {
		colmeia := domain.Colmeia{
			ID:           code,
			ColmeiaID:    intPtr(code),
			Species:      domain.PlebeiaEmerina,
			Status:       domain.Developing,
			StartingDate: mockTime,
		}
		colmeias = append(colmeias, colmeia)
	}

	for _, code := range codesForTB {
		colmeia := domain.Colmeia{
			ID:           code,
			ColmeiaID:    intPtr(code),
			Species:      domain.ScaptotrigonaBipunctata,
			Status:       domain.Developing,
			StartingDate: mockTime,
		}
		colmeias = append(colmeias, colmeia)
	}

	return colmeias
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prima101112/bababos-backend/migrations/seed"
	"github.com/prima101112/bababos-backend/pkg/handler"
	"github.com/prima101112/bababos-backend/repository"
)

func NewServer() {

	//initialize database repository
	db := repository.NewRepository()
	defer db.DB.Close()

	// initiate customer service
	customerService := handler.NewCustomerhandler(&repository.Repository{DB: db.DB})

	// initiate historicalpo service
	historicalpoService := handler.NewHistoricalpoHandler(&repository.Repository{DB: db.DB})

	// initiate historicalpo service
	pricelistsService := handler.NewPricelistsHandler(&repository.Repository{DB: db.DB})

	// initiate historicalpo service
	suggestionService := handler.NewSuggestionHandler(&repository.Repository{DB: db.DB})

	//cors
	r := mux.NewRouter()

	// Add CORS middleware
	r.Use(corsMiddleware)

	// Define routes
	r.HandleFunc("/customers", customerService.GetCustomersHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/historicalpo", historicalpoService.GetHistoricalposHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/pricelist", pricelistsService.GetPricelistsHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/suggestion", suggestionService.GetSuggestionHandler).Methods("GET", "OPTIONS")

	// Create a server and specify the address and port to listen on
	serverAddr := ":8080" // You can change the port as needed
	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	fmt.Printf("Server listening on %s...\n", serverAddr)

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func main() {
	//enable in dev mode always truncate before seeding
	// seeding not needed in staging or production
	seed.Seed()
	NewServer()
}

// Define the CORS middleware function
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "OK")
			return
		}

		next.ServeHTTP(w, r)
	})
}
